package main

import (
	"encoding/json"
	"fmt"
	"learner/pkg/sort/util"
)

/**
	对象排序
 */

type Person struct {
	Name string
	Age int
}


func main() {
	persions := []interface{}{
		&Person{
			Name: "AAA",
			Age: 55,
		},
		&Person{
			Name: "BBB",
			Age: 22,
		},
		&Person{
			Name: "CCC",
			Age: 0,
		},
		&Person{
			Name: "DDD",
			Age: 22,
		},
		&Person{
			Name: "EEE",
			Age: 11,
		},
	}

	util.DescSortBodyByFieldName(persions,"Age")
	bytes, _ := json.Marshal(persions)
	fmt.Println(string(bytes))


}