package main

import (
	"fmt"
	"sync"
)

func exampleRaceCondition() {
	numGR := 10

	k := make(map[int]int)
	k[1] = 12

	var wg sync.WaitGroup
	var i int
	for i = 0; i < numGR; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			k[i] = i
		}()
	}

	k[2] = 10

	wg.Wait()

	fmt.Printf("k = %#v\n", k)
}

func exampleNoRaceCondition() {
	numGR := 10

	var wg sync.WaitGroup
	var i int
	var aMutex sync.Mutex

	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < numGR; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			aMutex.Lock()
			k[j] = j
			aMutex.Unlock()
		}(i)
	}

	wg.Wait()
	k[2] = 10
	fmt.Printf("k = %#v\n", k)
}

func main() {
	//exampleRaceCondition()
	exampleNoRaceCondition()
}
