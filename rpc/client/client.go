package main

import (
	"fmt"
	"github.com/francisco-serrano/awesomeProject/rpc/sharedRPC"
	"log"
	"net/rpc"
)

func main() {
	addr := "localhost:8080"

	c, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	args := sharedRPC.MyFloats{16, -.5}
	var reply float64

	if err := c.Call("MyInterface.Multiply", args, &reply); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Reply (Multiply): %f\n", reply)

	if err := c.Call("MyInterface.Power", args, &reply); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Reply (Power): %f\n", reply)
}
