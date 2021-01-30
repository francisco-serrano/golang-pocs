package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	m  = sync.Mutex{}
	v1 int
)

func change(i int) {
	m.Lock()

	time.Sleep(time.Second)

	v1 += 1
	if v1%10 == 0 {
		v1 -= 10 * i
	}

	m.Unlock()
}

func read2() int {
	m.Lock()

	a := v1

	m.Unlock()

	return a
}

func exampleA() {
	var wg sync.WaitGroup
	fmt.Printf("%d ", read2())

	numGR := 21
	for i := 0; i < numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Printf("-> %d", read2())
		}(i)
	}

	wg.Wait()
	fmt.Printf("-> %d\n", read2())
}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

var password = secret{
	password: "myPassword",
}

func changeB(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

func showB(c *secret) string {
	c.RWM.RLock()
	defer c.RWM.RUnlock()

	fmt.Print("show")
	time.Sleep(3 * time.Second)

	return c.password
}

func showWithLockB(c *secret) string {
	c.RWM.Lock()
	defer c.RWM.Unlock()

	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)

	return c.password
}

func exampleB() {
	//showFunction := showB
	showFunction := showWithLockB

	fmt.Println("Pass:", showFunction(&password))

	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Go Pass:", showFunction(&password))
		}()
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		changeB(&password, "123456")
	}()

	wg.Wait()
	fmt.Println("Pass:", showFunction(&password))
}

type atomCounter struct {
	val int64
}

func (c *atomCounter) Add(v int64) {
	atomic.AddInt64(&c.val, v)
}

func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func exampleC() {
	X := 100 // goroutines
	Y := 200 // value

	var wg sync.WaitGroup
	var counter atomCounter

	for i := 0; i < X; i++ {
		wg.Add(1)
		go func(no int) {
			defer wg.Done()
			for i := 0; i < Y; i++ {
				counter.Add(1)
				//counter.val++
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(counter.Value())
}

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func exampleD() {
	n := 10

	fmt.Printf("Goint to create %d random numbers.\n", n)

	rand.Seed(time.Now().Unix())

	go monitor()

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * n))
		}()
	}

	wg.Wait()
	fmt.Printf("\nLast value: %d\n", read())
}

func main() {
	//exampleB()
	//exampleC()
	exampleD()
}
