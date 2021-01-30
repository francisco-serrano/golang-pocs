package main

import "fmt"

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

func main() {
	exampleA()
}
