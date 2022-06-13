package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/services"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := services.NewHelloServiceClient(conn)

	resp, err := client.SayHello(context.Background(), &services.HelloRequest{Req: "Hello gRPC"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Resp, resp.Timestamp.AsTime().Format("2006-01-02 15:04:05"))
}
