package main

import "fmt"

func main() {
	a := 3 + 2i
	b := 5 + 1i
	c := complex(5, 7)

	fmt.Println(a + b + c)
	fmt.Printf("%T\n", c)

	fmt.Println(a)
	fmt.Println(a - a)
	fmt.Println(a - 1)
	fmt.Printf("%T\n", c)
}
