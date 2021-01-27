package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/dev/random")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var seed int64
	binary.Read(f, binary.LittleEndian, &seed)

	fmt.Println("seed", seed)
}
