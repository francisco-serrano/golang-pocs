package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if _, err := io.WriteString(os.Stdout, scanner.Text()); err != nil {
			fmt.Println(err)
		}

		if _, err := io.WriteString(os.Stdout, "\n"); err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func main() {
	filename := ""

	arguments := os.Args
	if len(arguments) == 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			log.Fatal(err)
		}

		return
	}

	for i := 1; i < len(arguments); i++ {
		filename = arguments[i]
		if err := printFile(filename); err != nil {
			fmt.Println(err)
		}
	}
}