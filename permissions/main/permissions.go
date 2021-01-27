package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: permissions filename\n")
		return
	}

	filename := os.Args[1]

	info, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	mode := info.Mode()
	fmt.Println(filename, "mode is", mode.String()[1:10])
	fmt.Println(filename, "mode is", mode.String())
	fmt.Printf("%s mode is %d", filename, mode)

}
