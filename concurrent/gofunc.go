//理解 Go 里的函数
package main

import (
	"fmt"
	"time"
)

/**
	说明：

	1. go函数分为带名字函数和匿名函数
	2. 函数

		1. 形式参数列表描述了函数的参数名以及参数类型，这些参数作为局部变量，其值由函数调用者提供。
		2. 返回值列表描述了函数返回值的变量名以及类型，如果函数返回一个无名变量或者没有返回值.
		3. 返回值列表的括号是可以省略的。
	3. 匿名函数使用

		1. 定义变量名，是一个不难但是还费脑子的事情，对于那到只使用一次的函数，是没必要拥有姓名的。这才有了匿名函数。
 */
func main() {
	//fmt.Println("hello, world")

	go mygo("协程1号")
	go mygo("协程2号")
	time.Sleep(time.Second)

}

// 测试go routine的创建时间慢与main执行
/*func myTest(){
	fmt.Println("hello, go")
}*/

func mygo(name string)  {
	for i:=0; i<10 ;i++  {
		fmt.Printf("In gorountine %s\n",name)
		time.Sleep(10 * time.Millisecond)
	}
}