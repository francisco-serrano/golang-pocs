package main

import (
	"fmt"
	"log"
)

func panicFunction() {
	//panic("hello bitches")
	log.Panic("asdfasdfas")
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovery function", r)
	}

	fmt.Println("recovery teardown")
}

func main() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("recovering ohh yeah")
	//	}
	//
	//	fmt.Println("teardown")
	//}()

	defer recovery()

	panicFunction()

	fmt.Println("hello world")
}
