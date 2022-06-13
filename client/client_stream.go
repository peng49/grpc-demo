package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/services"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := services.NewHelloServiceClient(conn)
	stream, err := client.SayHelloForClientStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0;i<10;i++ {
		err = stream.Send(&services.HelloRequest{Req: "Hello Client Stream"})
		time.Sleep(time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	fmt.Println(res.Resp)
}
