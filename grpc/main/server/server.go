package main

import (
	"context"
	"fmt"
	"github.com/francisco-serrano/awesomeProject/grpc/main/pkg"
	"google.golang.org/grpc"
	"log"
	"net"
)

type MessageServer struct {

}

func (MessageServer) SayIt(ctx context.Context, r *pkg.Request) (*pkg.Response, error) {
	fmt.Println("Request Text:", r.Text)
	fmt.Println("Request SubText:", r.Subtext)

	response := &pkg.Response{
		Text:    r.Text,
		Subtext: "Got it!",
	}

	return response, nil
}

func main() {
	server := grpc.NewServer()

	pkg.RegisterMessageServiceServer(server, MessageServer{})

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Serving requets...")

	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}