package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		runtime.Breakpoint()
		fmt.Println("aaa")
	}(                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                )
	time.Sleep(1 * 1000)
	fmt.Println("end")
}
