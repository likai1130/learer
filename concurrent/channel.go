package main

import (
	"fmt"
	"sync"
	"time"
)

//场景1： 生产消费

// 生产者只能把数据往channel放
func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("正在生产第[%d]个\n", i)
	}
	close(ch)
	fmt.Println("colsed")
}

// 消费者只能把数据往channel取
func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("正在消费第[%d]个\n", n)
		time.Sleep(1 * time.Second)
	}
}

var wg sync.WaitGroup

func main() {
	proChan := make(chan int, 5)
	// 生产5个产品，并消费
	wg.Add(2)
	go func() {
		producer(proChan)
		wg.Done()
	}()
	go func() {
		consumer(proChan)
		wg.Done()
	}()
	wg.Wait()
}
