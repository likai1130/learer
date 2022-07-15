package main

import (
	"log"
	"sync"
	"time"
)

// 带缓冲channel 用法
// 消息队列；  FIFO
//用作计数信号量
var active = make(chan struct{},3)
var jobs = make(chan int,10)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- (i+1)
		}
		close(jobs)
	}()

	var wg  sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{}
			log.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<- active
			wg.Done()
		}(j)
	}
	wg.Wait()


}

