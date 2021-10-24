//信道
package main

import (
	"fmt"
	"time"
)

/**
信道，就是一个管道，连接多个goroutine程序 ，它是一种队列式的数据结构，遵循先入先出的规则。

有缓冲信道
无缓冲信道

*/

func main() {
	//singleChan()
		itemChan()
//	chanForlock()
}

/**
	5.  信道做锁
 */
func increment(ch chan bool, x *int) {
	ch <- true
	*x = *x + 1
	<- ch
}

func chanForlock() {
	pipline := make(chan bool, 1)

	var x int
	for i:=0; i<1000;i++{
		go increment(pipline, &x)
	}
	time.Sleep(time.Second)
	fmt.Println("X的值：",x)
}

/**
4. 遍历信道
*/

func ergodicChan(mychan chan int) {
	n := cap(mychan)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		mychan <- x
		x, y = y, x+y

	}
	close(mychan)
}

func itemChan() {
	pipline := make(chan int, 10)
	go ergodicChan(pipline)
	for k := range pipline {
		fmt.Println(k)
	}
}

/**

3. 单信道
<-chan 表示这个信道，只能从发出数据，对于程序来说就是只读
chan<- 表示这个信道，只能从外面接收数据，对于程序来说就是只写
*/
func singleChan() {
	type Receiver = <-chan int //只读
	type Sender = chan<- int   // 关键代码：定义别名类型  只写
	pipline := make(chan int)
	go func() {
		var sender Sender = pipline
		fmt.Println("准备发送数据 100")
		sender <- 100
	}()

	go func() {
		var receiver Receiver = pipline
		num := <-receiver
		fmt.Printf("接受到的数据 %d\n", num)
	}()
	time.Sleep(time.Second)
}

//1. 测试信道
func chanTest() {
	pipline := make(chan int)
	fmt.Printf("信道可缓冲 %d 个数据\n", cap(pipline))
	pipline <- 1
	a := <-pipline
	fmt.Println(a)
	fmt.Printf("信道当前有 %d 个数据\n", len(pipline))

}

//2. 双信道  无缓冲信道
func nocacheChan() {
	pipline := make(chan int)
	go func() {
		fmt.Println("send 100")
		pipline <- 100
	}()

	go func() {
		num := <-pipline
		fmt.Printf("接收到的数据是: %d", num)
	}()

	time.Sleep(time.Second)
}
