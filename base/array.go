package main

import (
	"fmt"
	"os"
)

/**
	数组
 */
func main() {
	InssertSeqlist()
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