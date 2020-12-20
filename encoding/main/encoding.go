package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//msg := "hello world"
	//
	//encoded := encode(msg)
	//decoded, err := decode(encoded)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(encoded)
	//fmt.Println(decoded)

	//Run()

	//RunWithSHA()

	RunWithHMAC()
}

func encode(msg string) string {
	return base64.URLEncoding.EncodeToString([]byte(msg))
}

func decode(msg string) (string, error) {
	b, err :=  base64.URLEncoding.DecodeString(msg)
	if err != nil {
		return "", fmt.Errorf("could not decode string %w", err)
	}

	return string(b), nil
}