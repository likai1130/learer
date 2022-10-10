package main

import (
	"fmt"
	"math/rand"
)


func main() {
	getRandom()
}

func getRandom() {
	fmt.Println(rand.Intn(100))
}
