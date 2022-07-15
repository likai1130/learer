package main

import (
	"fmt"
	"sync"
	"time"
)

// 用法二： 无缓冲信道1对n信号

func worker(i int)  {
	fmt.Printf("工作者[%d]: 正在工作\n", i)
	time.Sleep(1 *time.Second)
	fmt.Printf("工作者[%d]: 工作完毕\n", i)
}

type signal struct {}

func spawnGroup(f func(i int),num int,groupSignal <-chan signal) <- chan signal{
	c := make(chan  signal)
	var wg sync.WaitGroup

	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(i int) {
			<- groupSignal
			fmt.Printf("工作者 %d: 开始工作\n",i)
			f(i)
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

func main() {
	fmt.Println("开始一组工作者干活儿")
	groupSignal := make(chan signal)
	c := spawnGroup(worker, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("这组工作者开始干活了")
	close(groupSignal)
	<- c
	fmt.Println("这组工作者干完了")
}