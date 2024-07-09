package main

import (
	"net"
	gen_proto "proto-bug-demo/pkg/proto/api"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

	gen_proto.RegisterSwaggerServer(server, &Server{})
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
