package main

import "fmt"

type People struct {}

func (p *People) ShowA() {
	fmt.Println("show A")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("show B")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB()  {
	fmt.Println("teacher show B")
}

func main() {
	t:=Teacher{}
	t.ShowA()
}
