package main

import "fmt"

func printSomething() {
	for i := 0; i < 10; i++ {
		v := i * 2
		fmt.Println("v:", v)
		defer fmt.Println("deferred v:", v)
	}
}

func printAnotherThing() {
	var v int
	for i := 0; i < 10; i++ {
		v = i * 2
		fmt.Println("v:", v)
		defer fmt.Println("deferred v:", v)
	}
}

func main() {
	//printSomething()
	printAnotherThing()
}
