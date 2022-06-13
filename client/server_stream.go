package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/services"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := services.NewHelloServiceClient(conn)
	stream, err := client.SayHelloForServerStream(context.Background(), &services.HelloRequest{Req: "Hello Server Stream"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			//break
			log.Fatal(err)
		}

		fmt.Println(resp.Resp)
	}
}
