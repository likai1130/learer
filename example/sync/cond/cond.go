package main

import (
	"fmt"
	"sync"
	"time"
)

/**
	对mutex案例进行改进
 */
type signal struct {}

var ready bool

func worker(i int)  {
	fmt.Printf("worker %d: is working...\n",i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done\n",i)
}

func spawngroup(f func(i int), num int, groupSignal *sync.Cond) <- chan signal{
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			groupSignal.L.Lock()
			for !ready {
				groupSignal.Wait()
			}
			groupSignal.L.Unlock()
			fmt.Printf("worker %d: start to work... \n",i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	 go func() {
	 	wg.Wait()
	 	c <- signal(struct{}{})
	 }()
	return c
}

func main() {
	fmt.Println("start a group of workers...")
	greoupSignal := sync.NewCond(&sync.Mutex{})
	c := spawngroup(worker, 5, greoupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	greoupSignal.L.Lock()
	ready = true
	greoupSignal.Broadcast()
	greoupSignal.L.Unlock()
	<-c
	fmt.Println("the group of workers work done!")

}