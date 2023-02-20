package main

import (
	"fmt"
	"learner/example/reptile/tianjin/helper"
)

func main() {
	for i := 1; i <= 39; i++ {
		helper.ExampleScrape(i)
		fmt.Println()
	}
}
