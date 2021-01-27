package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	CLOSEA = false
	DATA   = make(map[int]bool)
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}

		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")

		if _, ok := DATA[x]; ok {
			CLOSEA = true
			continue
		}

		DATA[x] = true
		out <- x
	}

	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	sum := 0
	for x2 := range in {
		sum += x2
	}

	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	n1 := 1
	n2 := 999

	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		return
	}

	rand.Seed(time.Now().UnixNano())

	A := make(chan int)
	B := make(chan int)

	go first(n1, n2, A)
	go second(B, A)

	third(B)
}
