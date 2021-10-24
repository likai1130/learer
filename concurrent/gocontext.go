package main

import (
	"context"
	"fmt"
	"time"
)

/**
context 包

1. 方法：
	Deadline：
		返回第一个值是截止时间，context汇泽东触发cancel动作
		第二个值是一个布尔值，true表示截止时间，false表示没有设置截止时间，没有设置截止时间需要手动取消context
	Done：
		返回一个只读的通道（只有在被cancel后才会返回）类型为struct{},当这个通道只读意味着parent context已经发起了取消请求，根据这个信号，开发者手动做清理动作，取消goroutine

	Err: 返回context被取消的原因
	Value: 返回被绑定到Context的值，是一个键值对，所以要通过一个key才可以获取对应的值，这个值是线程安全的。


2. 为什么使用Context？

	当一个协程开启后，我们是无法强制关闭它的

	常见的关闭协程的原因有：

		1. goroutine自己跑完结束退出  （正常关闭）
		2. 主进程crash退出，goroutine被迫退出	（异常关闭，优化代码）
		3. 通过通道发出信号，导致协程的关闭( 设计手动关闭)

*/
func chainCloseTest() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case a := <-stop:
				fmt.Printf("监控退出，停止了...%v\n", a)
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(1 * time.Second)
			}

		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	time.Sleep(5 * time.Second)
}

func monitor(ch chan bool, number int) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("监控器%v，接收到通道值为：%v，监控结束\n", number, v)
			return
		default:
			fmt.Printf("监控器%v, 正在监控中... \n", number)
			time.Sleep(2 * time.Second)
		}
	}
}

func monitorTest() {
	stopSingal := make(chan bool)

	for i := 1; i <= 5; i++ {
		go monitor(stopSingal, i)
	}

	time.Sleep(1 * time.Second)
	close(stopSingal)

	//等待5s，若此时屏幕没有输出 <正在监控中> 就说明所有的goroutine都已经关闭
	time.Sleep(5 * time.Second)
	fmt.Println("主程序退出！")
}

//context 使用
func monitorContext(ctx context.Context, number int) {
	for {
		select {
		case v := <-ctx.Done():
			fmt.Printf("监控器%v，接收到通道值为：%v，监控结束\n", number, v)
			return
		default:
			fmt.Printf("监控器%v, 正在监控中... \n", number)
			time.Sleep(2 * time.Second)
		}
	}
}

func monitorContextTest() {
	//context.Background() 为 parent context 定义一个可取消的 context

	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= 5; i++ {
		go monitorContext(ctx, i)
	}
	time.Sleep(1 * time.Second)
	//关闭所有gorountine
	cancel()

	//等待5s，若此时屏幕没有输出 <正在监控中> 就说明所有的goroutine都已经关闭
	time.Sleep(5 * time.Second)
	fmt.Println("主程序退出！")

}

/**
context注意事项

1. 通常 Context 都是做为函数的第一个参数进行传递（规范性做法），并且变量名建议统一叫 ctx

2. Context 是线程安全的，可以放心地在多个 goroutine 中使用。

3. 当你把 Context 传递给多个 goroutine 使用时，只要执行一次 cancel 操作，所有的 goroutine 就可以收到 取消的信号

4. 不要把原本可以由函数参数来传递的变量，交给 Context 的 Value 来传递。

5. 当一个函数需要接收一个 Context 时，但是此时你还不知道要传递什么 Context 时，可以先用 context.TODO 来代替，而不要选择传递一个 nil。

6. 当一个 Context 被 cancel 时，继承自该 Context 的所有 子 Context 都会被 cancel。

*/
func main() {
	//chainCloseTest()
	//monitorTest()
	monitorContextTest()
}
