package main

import (
	"fmt"
	"os"
)

type xxx struct {
	Name string
}

func ( x xxx)ff()  {
	fmt.Println(x.Name)
}
/**
	数组
 */
func main() {
	//InssertSeqlist()
	//TArray()
	//aa()
	var aa xxx
	aa.ff()
}

/**
	[]int{1,2,3,4,5,6,7,10}
	8插入到10前面
 */
func InssertSeqlist() {
	data:=111
	i := 5
	//list := [...]int{1,2,3,4,5,6,7,8}
	list := make([]int,0,20)
	list = append(list, 1,2,3,4,5,6,7,8)
	if i <1 || i> len(list) + 1 {
		os.Exit(1)
	}

	list = append(list, 0)
	for j:= len(list);j>i ; j-- {
		list[j-1] = list[j-2]

		/*list[j] = list[j-1]*/
	}
	list[i-1] = data
	fmt.Println(list)
}

func TArray() {
	/* 创建切片 */
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)
	number2 := numbers[:2]
	printSlice(number2)
	number3 := numbers[2:5]
	printSlice(number3)
}
func printSlice(x []int){ fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)}

func aa() {
	for i := 3; i < 100; i++ {
		var j = 2
		for j = 2; j < i; j++ {
			if i%j == 0 {
				break
			}
		}
		if i == j {
			fmt.Print(i, " ")
		}
	}
}