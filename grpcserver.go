package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func server() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	grpcserver := grpc.NewServer()

	if grpcserver.Serve(lis); err != nil {
		log.Fatalf("Failed to server grpc server over port 9000: %v", err)

	}

}
