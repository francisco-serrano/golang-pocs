package main

import (
	"fmt"
	"log"
	"sync"
)

type sampleStruct struct {
	Message string `json:"message"`
}

func (s *sampleStruct) Reset() {
	s.Message = ""
}

var pool = sync.Pool{
	New: func() interface{} {
		return &sampleStruct{}
	}}

func main() {
	a := sampleStruct{Message: "hello world"}

	fmt.Println(a)

	pool.Put(&a)

	data, ok := pool.Get().(*sampleStruct)
	if !ok {
		log.Fatal("invalid assertion")
	}

	fmt.Println(*data)

	data.Reset()

	fmt.Println(*data)

	data, ok = pool.Get().(*sampleStruct)
	if !ok {
		log.Fatal("invalid assertion")
	}

	fmt.Println(*data)
}
