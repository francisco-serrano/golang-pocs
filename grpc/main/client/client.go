package main

import (
	"context"
	"fmt"
	"github.com/francisco-serrano/awesomeProject/grpc/main/pkg"
	"google.golang.org/grpc"
	"log"
)

func AboutToSayIt(ctx context.Context, m pkg.MessageServiceClient, text string) (*pkg.Response, error) {
	request := &pkg.Request{
		Text:    text,
		Subtext: "New Message!",
	}

	r, err := m.SayIt(ctx, request)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func main() {
	port := ":8080"

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pkg.NewMessageServiceClient(conn)

	r, err := AboutToSayIt(context.Background(), c, "My Message!")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response Text:", r.Text)
	fmt.Println("Response SubText:", r.Subtext)
}
