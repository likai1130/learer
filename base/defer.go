package main

import "fmt"

var name func() = C()

func A()  {
	fmt.Println("edasdasda")
}
func B(f func())  {

}
func C() func(){
	return A
}
func bar() int{
	defer func() int {
		fmt.Println("defer func in bar")
		return 2
	}()
	fmt.Println("call bar")
	fmt.Println("exit bar")
	return 1
}

func main() {
	i := bar()
	fmt.Println(i)
	name()
}