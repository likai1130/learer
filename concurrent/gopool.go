//协程池
package main

import (
	"fmt"
	"time"
)

type Pool struct {
	work chan func() //任务。无缓冲
	sem chan struct{} //数量 有缓冲
}

//创建协程池
func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem: make(chan struct{},size),
	}
}

//给协程池添加任务，使用 go worker 开启一个协程
func (p *Pool) NewTask(task func()){
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)

	}
}

//用于执行任务，for循环用于一直接收新的任务
func (p *Pool) worker(task func()) {
	defer func() {
		<- p.sem
	}()

	for {
		task()
		task = <- p.work
	}
}

func main() {
	pool := New(2)

	for i := 1; i <5 ; i ++{
		pool.NewTask(func() {
			time.Sleep(2 * time.Second)
			fmt.Println(time.Now())
		})
	}

	time.Sleep(5 * time.Second)
}