package main

import (
	"context"
	"fmt"

	"github.com/azzuwan/grpc-experiment/greetings"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:7070", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	fmt.Println("\n\nConnecting to client")

	client := greetings.NewGreeterClient(conn)
	req := greetings.HelloRequest{}
	res, err2 := client.SayHello(context.Background(), &req)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("\n\nRESPONSE ", res.Message)

}
