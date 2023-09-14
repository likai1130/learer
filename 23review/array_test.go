package _3review

import (
	"fmt"
	"testing"
)

/*
*
测试数组值拷贝
*/
func TestValueCopy(t *testing.T) {
	a := [3]int{1, 2, 3}
	fmt.Printf("a:%v, %p\n", a, &a)

	b := a
	fmt.Printf("b:%v, %p\n", b, &b)

	CopyArray(a)
}

func CopyArray(c [3]int) {
	fmt.Printf("c:%v, %p\n", c, &c)
}

/*
*
测试切片也是值拷贝
*/
func TestSliceCopy(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Printf("a: %p\n", a)

	b := a
	fmt.Printf("b: %p\n", b)
	CopySlice(a)
	fmt.Printf("a:%v %p\n", a, a)
	fmt.Printf("b: %p\n", b)
}

func CopySlice(c []int) {
	c = append(c, 333)
	fmt.Printf("c: %p\n", c)
}

func TestCopy(t *testing.T) {
	ints := make([]int, 0, 2)
	ints = append(ints, 10)
	ints = append(ints, 20)

	var a = []int{1, 2, 3}
	c := a
	fmt.Printf("ints: %v, %p\n", ints, ints)

	fmt.Printf("a: %v, %p\n", a, a)
	fmt.Printf("c: %v, %p\n", c, c)
	fmt.Println(copy(ints, a))

	fmt.Printf("copy ints: %v, %p\n", ints, ints)
	fmt.Printf("copy a: %v, %p\n", a, a)
	fmt.Printf("copy c: %v, %p\n", c, c)
}

func TestInitSlice(t *testing.T) {
	var a []int
	b := []int{}
	c := new([]int)
	d := make([]int, 0, 0)
	e := make([]int, 0)
	fmt.Printf("init a : %v, %p\n", a, &a)
	fmt.Printf("init b : %v, %p\n", b, &b)
	fmt.Printf("init c : %v, %p\n", c, &c)
	fmt.Printf("init d : %v, %p\n", d, &d)
	fmt.Printf("init e : %v, %p\n", e, &e)

	if a == nil {
		fmt.Printf("init new a : %v, %p\n", a, a)
		//fmt.Println("a is nil")
	}
}

func TestMap(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"d": 1,
		"c": 1,
		"b": 1,
	}

	fmt.Println(len(m))
}
