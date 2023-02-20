package main

import "fmt"

type S struct {
	Name string
}

var mapTest map[string]string

var sliceTest []int

func main() {
	//mapTest["aa"] = "aa" // assignment to entry in nil map
	//fmt.Println(mapTest["aa"]) // ok, 不能存，可以读

	mapTest = make(map[string]string)
	mapTest["aa"] = "aa"
	fmt.Println(mapTest["aa"])

	/*fmt.Println(sliceTest)           // []
	fmt.Println(sliceTest[0])        //panic: runtime error: index out of range [0] with length 0
	sliceTest = append(sliceTest, 1) //可正常使用*/

	sliceTest = make([]int, 10)
	fmt.Printf("%T\n", sliceTest) // []
	fmt.Println(sliceTest[0])

}
