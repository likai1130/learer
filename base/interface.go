package main

import "fmt"

/**
	interface 重新认识

	Go 语言中的函数调用都是值传递的，变量会在方法调用前进行类型转换。
 */

type Phonex interface {
	call()
}

type iPhone struct {
	name string
}

func (phone iPhone) call()  {
	fmt.Println("Hello, iPhone.")
}

func (phone iPhone) send_wechat()  {
	fmt.Println("Hello, Wechat.")
}


func main() {
	phone := iPhone{name:"ming's iphone"}
	phone.call()
	phone.send_wechat()
}
