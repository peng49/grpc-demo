package main

import (
	"google.golang.org/grpc"
	"grpc-demo/services"
	"log"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterHelloServiceServer(rpcServer, new(services.HelloService))

	lst, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}

	err = rpcServer.Serve(lst)
	if err != nil {
		log.Fatal(err)
	}
}
