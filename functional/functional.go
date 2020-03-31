package functional

import "fmt"

func Run() {
	fruits := []string{"orange", "apple", "banana", "grape"}

	out := mapElements(fruits, func(s string) int {
		return len(s)
	})

	fmt.Println(out)

	addFunc1 := add(10)
	addFunc2 := add(20)
	addFunc3 := add(30)

	fmt.Println(addFunc1(15))
	fmt.Println(addFunc2(15))
	fmt.Println(addFunc3(15))
}

func mapElements(arr []string, fn func(it string) int) []int {
	var newArray []int
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}

	return newArray
}

func add(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}
