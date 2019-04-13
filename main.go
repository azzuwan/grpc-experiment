package main

import (
	"fmt"
	"net"

	"github.com/azzuwan/grpc-experiment/greetings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":7070")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	greetings.RegisterGreeterServer(srv, &greetings.GServer{})
	reflection.Register(srv)
	fmt.Println("\n\n Server Started")
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}

}
