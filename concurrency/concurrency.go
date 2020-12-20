package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)
	}()

	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}

		close(out)
	}()

	return out
}

func merge(cs ...chan int) chan int {
	var wg sync.WaitGroup

	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}

		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func genWithContext(ctx context.Context) chan int {
	out := make(chan int)

	go func() {
		generateValues := true
		timeout := ctx.Done()
		accum := 0

		for generateValues {
			select {
			case out <- accum:
				accum++
				time.Sleep(500 * time.Millisecond)
			case <-timeout:
				generateValues = false
			}
		}

		close(out)
	}()

	return out
}

func genWithWaitGroup(n int) chan int {
	wg := sync.WaitGroup{}
	wg.Add(n)

	c := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			c <- i
			time.Sleep(200 * time.Millisecond)
			wg.Done()
		}

		close(c)
	}()

	wg.Wait()

	return c
}

func multiplyBy2(c chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range c {
			out <- n * 2
		}

		close(out)
	}()

	return out
}

func Run() {
	for n := range multiplyBy2(genWithWaitGroup(10)) {
		fmt.Println(n)
	}

	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//
	//for n := range sq(genWithContext(ctx)) {
	//	fmt.Println(n)
	//}

	//for n := range sq(gen(1, 2, 3, 4, 5)) {
	//	fmt.Println(n)
	//
	//	time.Sleep(1 * time.Second)
	//}

	//in := gen(1, 2, 3, 4, 5, 6)
	//
	//c1 := sq(in)
	//c2 := sq(in)
	//
	//for n := range merge(c1, c2) {
	//	fmt.Println(n)
	//}
}
