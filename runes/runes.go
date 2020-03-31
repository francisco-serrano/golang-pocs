package runes

import (
	"fmt"
	"reflect"
	"unicode"
)

func Run() {
	rune1 := 'B'
	rune2 := 'g'
	rune3 := '\a'

	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s\n", rune1, rune1, reflect.TypeOf(rune1))
	fmt.Printf("Rune 2: %c; Unicode: %U; Type: %s\n", rune2, rune2, reflect.TypeOf(rune2))
	fmt.Printf("Rune 3: %c; Unicode: %U; Type: %s", rune3, rune3, reflect.TypeOf(rune3))

	slc := []rune{'♠', '♠', '♧', '♡', '♬'}

	for i, v := range slc {
		fmt.Printf("Character: %c, Unicode: %U, Position: %d\n", unicode.ToUpper(v), v, i)
	}
}
