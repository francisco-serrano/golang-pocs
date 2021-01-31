package main

import (
	"fmt"
	"math/rand"
	"os"
)

func fib1(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fib2(n-1) + fib2(n-2)
}

func fib3(n int) int {
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

	return fn[n]
}

func exampleA() {
	fmt.Println(fib1(40))
	fmt.Println(fib2(40))
	fmt.Println(fib3(40))
}

var (
	bufferSize int
	fileSize   int
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createBuffer(buf *[]byte, count int) {
	*buf = make([]byte, count)
	if count == 0 {
		return
	}

	for i := 0; i < count; i++ {
		intByte := byte(random(0, 100))
		if len(*buf) > count {
			return
		}

		*buf = append(*buf, intByte)
	}
}

func Create(dst string, b, f int) error {
	if _, err := os.Stat(dst); err != nil {
		return fmt.Errorf("File %s already exists.", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer func() {
		if err := destination.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	buf := make([]byte, 0)
	for {
		createBuffer(&buf, b)
		buf = buf[:b]

		if _, err := destination.Write(buf); err != nil {
			return err
		}

		if f < 0 {
			break
		}

		f -= len(buf)
	}

	return nil
}

func exampleB() {
	output := "/tmp/randomFile"

	bufferSize = 10
	fileSize = 100000

	if err := Create(output, bufferSize, fileSize); err != nil {
		fmt.Println(err)
	}

	if err := os.Remove(output); err != nil {
		fmt.Println(err)
	}
}

func main() {
	//exampleA()
	exampleB()
}
