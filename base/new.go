package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	mutex.Lock()

	rwMutex := sync.RWMutex{}
	rwMutex.Lock()
	rwMutex.RLock()

	s := make([]int, 1024, 1024)

	s = append(s, 1)

	fmt.Println(len(s), cap(s))

	for i := 1; i <= 512; i++ {
		s = append(s, i)

	}
	fmt.Println(len(s), cap(s))
}
