package main

import (
	"fmt"
)

type student struct {
	Name	string
	Age 	int
}
/**
	for循环
 */
func main() {
	m := make(map[string]*student)
	stus := []student{
		{"zhang3", 18},
		{"li4", 19},
		{"wang5", 20},
	}

	for i, stu := range stus {
		m[stu.Name] = &stus[i]
	}

	for _, stu := range m {
		fmt.Println(stu.Name, stu.Age)
	}
}

