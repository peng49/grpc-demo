package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/services"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := services.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	stream, err := client.SayHelloForStream(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//响应
	go func() {
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
	}()

	//调用
	for i := 0; i < 10; i++ {
		err = stream.Send(&services.HelloRequest{Req: "Hello Stream"})
		time.Sleep(time.Second)
	}

	err = stream.CloseSend()
	if err != nil {
		fmt.Println(err)
	}
}
