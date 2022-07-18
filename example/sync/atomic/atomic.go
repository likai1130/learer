package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// atomic 计数 == mutex.lock count++ mutex.unlock
var count int64

func worker(i int)  {
	atomic.AddInt64(&count,1)
}

func exec(f func(i int),num int)  {
	wg := sync.WaitGroup{}
	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(i int) {
			f(i)
			wg.Done()
			fmt.Printf("第%d个计算结果%d\n",i,count)
		}(i)
	}
	wg.Wait()
	fmt.Println(count)

}

func main() {
	exec(worker,200)
}