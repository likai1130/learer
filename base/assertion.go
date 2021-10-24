package main

import "fmt"

/**
	断言Type Assertion
	检查 i 是否为 nil

	检查 i 存储的值是否为某个类型

	1. 类型断言，仅能对静态类型为空接口（interface{}）的对象进行断言，否则会抛出错误，
	2. 类型断言完成后，实际上会返回静态类型为你断言的类型的对象，而要清楚原来的静态类型为空接口类型（interface{}），这是 Go 的隐式转换。

 */
func main() {
	var i interface{} = 10

	t1,ok := i.(int)
	fmt.Printf("%d-%t\n", t1, ok)

	t2,ok := i.(string)
	fmt.Printf("%s-%t\n", t2, ok)

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok)

	k = 10
	t4, ok := k.(interface{})
	fmt.Printf("%d-%t\n", t4, ok)

	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok)

	findType("da")
	findType(11)
	var l interface{}
	findType(l)

	findType(10.23)
}


/**
	Type Switch
 */

func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x,"is int")
	case string:
		fmt.Println(x,"is string")
	case nil:
		fmt.Println(x,"is nil")
	default:
		fmt.Println(x,"not type matched")
	}
}
