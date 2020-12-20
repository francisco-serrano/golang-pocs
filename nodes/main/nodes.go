package main

import "fmt"

func functionOne(x int) {
	fmt.Println(x)
}

func main() {
	varOne := 1
	varTwo := 2

	fmt.Println("hello there!")

	functionOne(varOne)
	functionOne(varTwo)
}
