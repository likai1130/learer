package main

import (
	"fmt"
	"strconv"
)

func main() {
	chainId := "0x4"
	if chainId[:2] != "0x" {
		fmt.Println(false)
	}
	fmt.Printf("%x\n",42)

	fmt.Println(fmt.Sprintf("%d","3e"))

	n, err := strconv.ParseUint("2a", 16, 32)
	fmt.Println(n,err)

}

