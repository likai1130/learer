package main

import (
	"fmt"
	"math"
)

// 返回一个“返回int的函数”
func fibonacci() func() int {
	a, b, tmp := 0,1,0
	return func() int{
		tmp = a
		a = b
		b = tmp + b
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	vertex := Vertex{}
	abs := Vertex.Abs(vertex)
	fmt.Println(abs)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}




