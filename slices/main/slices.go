package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 6)

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	fmt.Println()

	b := make([]int, 1000000)

	fmt.Println(len(b))
	fmt.Println(cap(b))

	b = append(b, 1)

	fmt.Println(len(b))
	fmt.Println(cap(b))

	b = append(b, 1)

	fmt.Println(len(b))
	fmt.Println(cap(b))

	c := []int{1, 2, 3, 4}
	d := c[0:2]

	fmt.Println(c)
	fmt.Println(d)

	e := [5]int{1, 2, 3, 4, 5}

	fmt.Println(e)

	f := make([]int, 10)
	f = append(f, 10)

	g := new([]int)

	fmt.Println(f, len(f), cap(f), f[11])
	fmt.Println(g, len(*g), cap(*g))
}
