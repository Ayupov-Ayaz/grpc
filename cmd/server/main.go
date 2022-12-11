package main

import (
	"fmt"
	"github.com/ayupov-ayaz/grpc/pkg/adder"
	"github.com/ayupov-ayaz/grpc/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func run() error {
	s := grpc.NewServer()
	srv := adder.NewGRPCServer()
	api.RegisterAdderServer(s, srv)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	fmt.Println("server is running...")

	if err := s.Serve(listener); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
