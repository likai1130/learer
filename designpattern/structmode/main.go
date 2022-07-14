package main

import (
	"fmt"
)

func main() {
	var num1 int
	var num2 int
	fmt.Println("请输出第1个数字:")
	fmt.Scanf("%d", &num1)
	if num1 <= 0 {
		fmt.Println("第1个数字不正确,请重新输入!")
		return
	}
	fmt.Println("请输出第2个数字:")
	fmt.Scanf("%d", &num2)
	if num2 <= 0 {
		fmt.Println("第2个数字不正确,请重新输入!")
		return
	}
	fmt.Printf("运行结果:%d\n", num1+num2)
	fmt.Println("结束!")
}
