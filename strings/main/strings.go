package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func exampleA() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)
	fmt.Printf("x: %x\n", sLiteral)

	fmt.Printf("sLiteral length: %d\n", len(sLiteral))

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
	}
	fmt.Println()

	fmt.Printf("q: %q\n", sLiteral)
	fmt.Printf("+q: %+q\n", sLiteral)
	fmt.Printf(" x: % x\n", sLiteral)

	fmt.Printf("s: As a string: %s\n", sLiteral)
}

func exampleB() {
	r := strings.NewReader("test")

	fmt.Println("r length:", r.Len())

	b := make([]byte, 1)

	for {
		n, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println(err)
			continue
		}

		fmt.Printf("read %s bytes: %d\n", b, n)
	}

	s := strings.NewReader("This is an error!\n")

	fmt.Println("r length: ", s.Len())

	n, err := s.WriteTo(os.Stderr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("wrote %d bytes to os.stderr\n", n)
}

func main() {
	//exampleA()
	exampleB()
}
