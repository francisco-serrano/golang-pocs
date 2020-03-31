package shift

import "fmt"

func Run() {
	var a uint8 = 192

	fmt.Println(a)

	fmt.Println(a << 1)

	fmt.Println(a >> 4)

	fmt.Println("----------------")

	var b int8 = 0

	fmt.Println(b)

	fmt.Println(b << 1)
}
