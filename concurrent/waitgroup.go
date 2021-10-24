//waitgrioup 优雅解决协程退出
package main

import (
	"fmt"
	"sync"
)

/**
	不要通过共享内存来通信，要通过通信来共享内存

	WaitGroup:

	 Add: 初始值为0，传入的值会往计数器上增加。这里直接传入子协程的数量
     Done: 当某个子协程完成后，可调用此方法，会从计数器上减一，通常可以使用defer调用
	 wait: 阻塞当前协程，直到实例里的计数器归零。
 */


func main() {
	waitGroupTest()
}

func waitGroupTest()  {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(1,&wg)
	go worker(2,&wg)
	wg.Wait()
}
func worker(x int, wg *sync.WaitGroup)  {
	defer wg.Done()
	for i := 0; i< 5 ; i++  {
		fmt.Printf("worker %d: %d \n",x,i)
	}
}
