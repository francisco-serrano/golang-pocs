package main

import (
	"fmt"
	"github.com/francisco-serrano/awesomeProject/rpc/sharedRPC"
	"log"
	"math"
	"net"
	"net/rpc"
)

type MyInterface struct {
}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func (m *MyInterface) Multiply(args *sharedRPC.MyFloats, reply *float64) error {
	*reply = args.A1 * args.A2
	return nil
}

func (m *MyInterface) Power(args *sharedRPC.MyFloats, reply *float64) error {
	*reply = Power(args.A1, args.A2)
	return nil
}

func main() {
	port := ":8080"

	myInterface := new(MyInterface)

	rpc.Register(myInterface)

	t, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}
}