package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var buffer bytes.Buffer

	if _, err := buffer.Write([]byte("This is")); err != nil {
		log.Fatal(err)
	}

	if _, err := fmt.Fprintf(&buffer, " a string!\n"); err != nil {
		log.Fatal(err)
	}

	if _, err := buffer.WriteTo(os.Stdout); err != nil {
		log.Fatal(err)
	}

	if _, err := buffer.WriteTo(os.Stdout); err != nil {
		log.Fatal(err)
	}

	buffer.Reset()

	buffer.Write([]byte("Mastering Go!"))

	r := bytes.NewReader(buffer.Bytes())

	fmt.Println(buffer.String())

	for {
		b := make([]byte, 3)

		n, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println(err)
			continue
		}

		fmt.Printf("Read %s bytes: %d\n", b, n)
	}
}
