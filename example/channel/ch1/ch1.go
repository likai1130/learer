package main

import (
	"fmt"
	"time"
)

// 无缓冲信道惯用方法
//第一种用法：用作信号传递 1:1

type signal struct {}

func worker()  {
	println("正在work")
	time.Sleep(10 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("开始work")
		f()
		c <- signal{}
	}()
	return c
}

func main() {
	println("创建一个工作者")
	c:= spawn(worker)
	<-c
	fmt.Println("工作者干完了")
}