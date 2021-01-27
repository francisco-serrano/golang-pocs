package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
)

func open() {
	os.NewFile(1, "stdout")
}

func exampleA() {
	open()

	runtime.GC()

	if _, err := fmt.Println("some text"); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("could not print the text: %w", err))
	}
}

func exampleB() {
	arguments := os.Args
	if len(arguments) != 3 {
		log.Fatal("<buffer size> <filename>")
	}

	bufferSize, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Fatal(err)
	}

	file := arguments[2]
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for {
		readData := readSize(f, bufferSize)
		if readData != nil {
			fmt.Print(string(readData))
		} else {
			break
		}
	}
}

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)

	n, err := f.Read(buffer)
	if err != nil {
		if err == io.EOF {
			return nil
		}

		log.Fatal(err)
	}

	return buffer[0:n]
}

func main() {
	//exampleA()
	exampleB()
}
