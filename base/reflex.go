package main

import (
	"fmt"
	"reflect"
)

/**
	反射

	核心 Type和Value

		reflect.Type

		reflect.Value
	三大定律：

		1. 反射可以将接口类型变量转换成为"反射类型对象"
		2. 反射可以将"反射类型对象"转换为接口类型变量
		3. 如果要修改"反射类型对象"其类型必须是可写的

 */




 /**
 	定律一

		为了实现从接口变量到反射对象的转换，需要提到 reflect 包里很重要的两个方法：

 		reflect.TypeOf(i) ：获得接口值的类型

		reflect.ValueOf(i)：获得接口值的值

  */
func low1() {
	var age interface{} =25
	fmt.Printf("原始接口变量类型为%T,值为%v\n",age,age)

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)
}


/**
	定律二

	从反射对象到接口变量的转换。

 	reflect.Value 的结构体会接收 Interface 方法，
	返回了一个 interface{} 类型的变量
	（注意：只有 Value 才能逆向转换，而 Type 则不行，这也很容易理解，如果 Type 能逆向，那么逆向成什么呢？）
 */
func low2() {
	var age interface{} =25
	fmt.Printf("原始接口变量类型为%T,值为%v\n",age,age)

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

	i := v.Interface().(int)

	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i)

}

/**
	定律三

		 settable （可设置性，或可写性）的概念。

	以内golang里面函数传递是值传递，只要你传递的不是变量的指针，你在函数内部对变量的修改是不会影响到原始的变量的。

	回到反射上来，当你使用 reflect.Typeof 和 reflect.Valueof 的时候，如果传递的不是接口变量的指针，
反射世界里的变量值始终将只是真实世界里的一个拷贝，你对该反射对象进行修改，并不能反映到真实世界里。

	- 不是接收变量指针创建的反射对象，不具备可写性
	- 是否具备可写性，使用CanSet来获取得知
	- 对不具备可写性的对象进行修改，没有意义，认为是不可发会报错

注意：

	1. 创建反射对象时传入变量的指针
	2. 使用Elem()函数返回指针指向的数据
 */
func low3()  {
	var name string = "kli"
	v1 := reflect.ValueOf(&name)
	fmt.Println("v1 可写性为：",v1.CanSet())

	v2 := v1.Elem()
	fmt.Println("v2 可写性为:", v2.CanSet())
	v2.SetString("likai")
	fmt.Println("通过反射对象进行更新后，真实世界里 name 变为：", name)

}

func main() {

	//low2()
	low3()

}

