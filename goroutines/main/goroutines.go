package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func exampleA() {
	count := 20

	fmt.Printf("Going to create %d goroutines.\n", count)

	var wg sync.WaitGroup

	fmt.Printf("%#v\n", wg)

	for i := 0; i < count; i++ {
		wg.Add(1)

		go func(x int) {
			defer wg.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("%#v\n", wg)

	wg.Wait()

	fmt.Println("\nExiting...")
}

func exampleB() {
	c := make(chan int)

	go writeToChannel(c, 10)

	time.Sleep(1 * time.Second)

	fmt.Println("Read:", <-c)

	time.Sleep(1 * time.Second)

	if _, ok := <-c; ok {
		fmt.Println("Channel is open!")
		return
	}

	fmt.Println("Channel is closed!")
}

func writeToChannel(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func exampleC() {
	willClose := make(chan int, 10)

	willClose <- -1
	willClose <- 0
	willClose <- 2

	<-willClose
	<-willClose
	<-willClose

	close(willClose)

	read, ok := <-willClose

	fmt.Println(read, ok)
}

func f1(c chan int, x int) {
	fmt.Println(x)
	c <- x
}

func f2(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
}

func f3(out <-chan int, in chan<- int) {
	x := <-out
	fmt.Println(x)
	in <- x
}

func exampleD() {
	var i byte
	fmt.Println(i)
	fmt.Println(i+1)
	go func() {
		for i := 0; i < 255; i++ {

		}
		fmt.Println("AAAA")
	}()

	fmt.Println("Leaving goroutine!")

	runtime.Gosched()
	runtime.GC()

	fmt.Println("Good bye!")
}

func main() {
	//exampleA()
	//exampleB()
	//exampleC()
	exampleD()
}
