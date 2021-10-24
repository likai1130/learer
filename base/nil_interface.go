package main

import "fmt"

/**
	空接口

	定义:

		空接口是特殊形式的接口类型，普通的接口都有方法，而空接口没有定义任何方法口，也因此，我们可以说所有类型都至少实现了空接口。

	特点：

		每一个接口都包含两个属性，一个是值，一个是类型。对于空接口来说，这两者都是 nil

*/

type empty_interface interface {


}

func main() {
}

/**
	场景一

		通常我们会直接使用 interface{} 作为类型声明一个实例，而这个实例可以承载任意类型的值。

 */
func scene() {

	var i interface{}

	i = 1
	fmt.Println(i)

	i = "hello"
	fmt.Println(i)

	i = false
	fmt.Println(i)
}

/**
	场景二

	如果想让你的函数可以接收任意类型的值 ，也可以使用空接口
 */

func myFunc(iface interface{}) {
	fmt.Println(iface)
}
func myFuncUpgrade(ifaces ...interface{}) {
	for _,iface := range ifaces  {
		fmt.Println(iface)
	}
}
func myFuncTest() {
	a := 10
	b := "hello"
	c := true

	myFunc(a)
	myFunc(b)
	myFunc(c)
	myFuncUpgrade(a,b,c)
}

/**
	场景三

		你也定义一个可以接收任意类型的 array、slice、map、strcut，例如这边定义一个切片
 */
func mySliceInterface()  {

	any := make([]interface{}, 5)
	any[0] = 11
	any[1] = "hello world"
	any[2] = []int{11, 22, 33, 44}
	for _, value := range any {
		fmt.Println(value)
	}

}
