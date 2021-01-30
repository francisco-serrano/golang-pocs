package main

import (
	"fmt"
	"github.com/pkg/profile"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func fib1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}

	time.Sleep(time.Millisecond)

	return int64(fib2(n-1)) + int64(fib2(n-2))
}

func fib2(n int) int {
	fn := make(map[int]int)

	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}

		fn[i] = f
	}

	time.Sleep(50 * time.Millisecond)

	return fn[n]
}

func N1(n int) bool {
	k := math.Floor(float64(n/2 + 1))
	for i := 2; i < int(k); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func N2(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func exampleA() {
	cpuFile, err := os.Create("/tmp/cpuProfile.out")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := cpuFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal(err)
	}

	defer pprof.StopCPUProfile()

	total := 0
	for i := 0; i < 100000; i++ {
		if n := N1(i); n {
			total += 1
		}
	}
	fmt.Println("Total primes:", total)

	total = 0
	for i := 2; i < 100000; i++ {
		if n := N2(i); n {
			total += 1
		}
	}
	fmt.Println("Total primes:", total)

	for i := 0; i < 90; i++ {
		n := fib1(i)
		fmt.Print(n, " ")
	}
	fmt.Println()
	for i := 0; i < 90; i++ {
		n := fib2(i)
		fmt.Print(n, " ")
	}
	fmt.Println()

	runtime.GC()

	memory, err := os.Create("/tmp/memoryProfile.out")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := memory.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for i := 0; i < 10; i++ {
		if s := make([]byte, 5000000); s == nil {
			fmt.Println("Operation failed!")
		}

		time.Sleep(50 * time.Millisecond)
	}

	if err := pprof.WriteHeapProfile(memory); err != nil {
		log.Fatal(err)
	}
}

var VARIABLE int

func exampleB() {
	defer profile.Start(profile.MemProfile).Stop()

	total := 0
	for i := 2; i < 200000; i++ {
		if n := N3(i); n {
			total++
		}
	}
	fmt.Println("Total:", total)

	total = 0
	for i := 0; i < 5000; i++ {
		for j := 0; j < 400; j++ {
			k := Multiply(i, j)
			VARIABLE = k
			total++
		}
	}
	fmt.Println("Total:", total)
}

func N3(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func Multiply(a, b int) int {
	if a == 1 {
		return b
	}

	if a == 0 || b == 0 {
		return 0
	}

	if a < 0 {
		return -Multiply(-a, b)
	}

	return b + Multiply(a-1, b)
}

func main() {
	exampleA()
	//exampleB()
}
