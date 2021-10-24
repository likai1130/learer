package main

import (
	"fmt"
	"sort"
)

/*
	对结构按多字段排序
*/
type student struct {
	Name string
	Age  int
}

type stus []student
func(s stus) Len() int { return len(s) }
func(s stus) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type sortByName struct{ stus }

// 按名字排序
func(m sortByName) Less(i, j int) bool {
	return m.stus[i].Name > m.stus[j].Name
}


type sortByAge struct { stus }

// 按年龄排序
func(m sortByAge) Less(i, j int) bool {
	return m.stus[i].Age > m.stus[j].Age
}


func main() {
	s := stus{
		{
			Name: "test123",
			Age:  20,
		},
		{
			Name: "你",
			Age:  22,
		},
		{
			Name: "xxx",
			Age:  21,
		},
		{
			Name: "xxx",
			Age:  55,
		},
		{
			Name: "xxx",
			Age:  25,
		},
		{
			Name: "xxx",
			Age:  88,
		},
	}

	sort.Stable(sortByAge{s})
	//sort.Stable(sortByName{s})
	fmt.Println(s)

}