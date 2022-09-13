package main

import "sort"

func main() {
	for i:=1;i<=100;i++ {
		mt()
	}
}
const N = 3

func mt() {
	keys := []int{}
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		m[i] = i
		keys = append(keys, i)
	}

	sort.Ints(keys)
	for k := range keys {
		print(m[k])
	}
}
