package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	arguments := []int{5, 3, 1, 6, 7, 4, 2}

	var wg sync.WaitGroup
	for _, arg := range	arguments {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Print(n, " ")
		}(arg)
	}

	wg.Wait()
	fmt.Println()
}
