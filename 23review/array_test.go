package main

import (
	"fmt"
	"testing"
)

func TestValueCopy(t *testing.T) {
	a := [2]int{123}
	fmt.Printf("a: %p", &a)
}
