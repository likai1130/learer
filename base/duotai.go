package main

import (
"fmt"
	"unsafe"
)

type BaseIntf interface {
	f()
}


type D struct {
	x int
}

func (this *D)f()  {
	this.x = 200
}


func main() {
	var d D
	d.x = 100
	/*var b BaseIntf = &d*/

	fmt.Printf("%T, %v\n", d, d)

	d.f()
	fmt.Printf("%T, %v\n", d ,d)
	sizeof := unsafe.Sizeof(d)
	fmt.Println(sizeof)

	// b.x = 1   // b.x undefined (type BaseIntf has no field or method x)
}
