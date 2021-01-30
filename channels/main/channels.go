package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(4 * time.Second):
			fmt.Println("\ntime.After()!")
		}
	}
}

func exampleA() {
	rand.Seed(time.Now().Unix())

	createNumber := make(chan int)
	end := make(chan bool)

	n := 10

	fmt.Printf("Going to create %d random numbers.\n", n)

	go gen(0, 2*n, createNumber, end)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Exiting...")

	end <- true
}

func exampleB() {
	c1 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "c1 OK"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout c1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "c2 OK"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(4 * time.Second):
		fmt.Println("timeout c2")
	}
}

func exampleC() {
	var wg sync.WaitGroup
	wg.Add(1)

	t := 6000

	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)

	if timeout(&wg, duration) {
		fmt.Println("Timed out!")
		return
	}

	fmt.Println("OK!")
}

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)

	go func() {
		defer close(temp)
		time.Sleep(5 * time.Second)

		w.Wait()
	}()

	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func exampleD() {
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	for i := 0; i < counter+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("Nothing more to be done!")
			break
		}
	}
}

func exampleE() {
	c := make(chan int)
	go add(c)
	go send(c)

	time.Sleep(3 * time.Second)
}

func add(c chan int) {
	sum := 0
	fmt.Println(time.Now().Format(time.RFC3339Nano))
	t := time.NewTimer(time.Second)
	fmt.Println(time.Now().Format(time.RFC3339Nano))

	for {
		select {
		case input := <-c:
			sum += input
		case v := <-t.C:
			c = nil
			fmt.Println(sum, v, v.Format(time.RFC3339Nano))
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func exampleF() {
	times := 5

	cc := make(chan chan int)

	for i := 0; i < times+1; i++ {
		f := make(chan bool)

		go f1(cc, f)

		ch := <-cc
		ch <- i

		for sum := range ch {
			fmt.Print("Sum(", i, ")=", sum)
		}

		fmt.Println()

		time.Sleep(time.Second)

		close(f)
	}
}

func f1(cc chan chan int, f chan bool) {
	c := make(chan int)
	cc <- c
	defer close(c)

	sum := 0
	select {
	case x := <-c:
		for i := 0; i < x; i++ {
			sum += i
		}

		c <- sum
	case <-f:
		return
	}
}

func exampleG() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go C(z)
	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)

	close(x)
	time.Sleep(3 * time.Second)
}

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("C()!")
}

func main() {
	//exampleA()
	//exampleB()
	//exampleC()
	//exampleD()
	//exampleE()
	//exampleF()
	exampleG()
}
