//锁机制
package main

import (
	"fmt"
	"sync"
	"time"
)

/**
并发场景中，优先使用chan，如果信道解决不了，则考虑锁机制

Mutex： 互斥锁
RWMutet： 读写锁
*/
func main() {
	//mutexTest()
	rwMutexTest()
}

/**
2. 读写锁 RWMutex

将资源的访问分为读操作和写操作

特点：
	1. 为了保证数据安全，规定了当有人还在读取数据时（读锁占用）时，不允许有人更新这个数据（写锁会阻塞）
	2. 为了保证程序的效率，多个人（线程）读取数据（拥有读锁）时，互不影响不会造成阻塞，它不会像Mutex那样只允许有一个人（线程）读取同一个数据

定义：
	var lock *sync.RWMutex
	lock = new(sync.RWMutex)

	lock:=&sync.RWMutex{}

	读锁： 用RLock开启锁。RUnlock释放锁
	写锁： 用Lock开启锁，Unlock释放锁

*/

func rwMutexTest() {
	lock := &sync.RWMutex{}
	lock.Lock()

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("第 %d 个协程准备开始.... \n", i)
			lock.RLock()
			fmt.Printf("第 %d 个协程获得读锁, sleep 1s 后，释放读锁 \n", i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}

	time.Sleep(time.Second * 2)
	fmt.Println("准备释放写锁，读锁不再阻塞")

	//写锁释放，读锁自由？
	lock.Unlock()
	fmt.Println("写锁已释放")
	//由于会等到读锁全部释放，才能获得写锁
	//四个协程全部完成继续执行
	lock.Lock()
	fmt.Println("程序退出")
	lock.Unlock()

}

/**
1. 互斥锁： Mutex
使用互斥锁，是为了保护一个资源不会因为并发操作而引起冲突导致数据不准确

定义：

var lock *sync.Mutex
lock = new(sync.Mutex)

lock := &sync.Mutex{}

注意：

	1. 同一协程里面，不要在尚未解锁时，再次使用加锁
	2. 同一协程里面，不要对已解锁的锁。再次解锁
	3. 加了锁以后，不要忘记解锁，必要时使用defer语句解锁

缺点:
	时间上浪费，导致程序性能低下
*/
func mutexTest() {
	var wg sync.WaitGroup
	lock := &sync.Mutex{}
	count := 0
	wg.Add(3)
	go add(&count, &wg, lock)
	go add(&count, &wg, lock)
	go add(&count, &wg, lock)
	wg.Wait()
	fmt.Println("count值为：", count)

}

func add(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count = *count + 1
		lock.Unlock()
	}
	defer wg.Done()
}
