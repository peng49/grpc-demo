package services

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"time"
)

type HelloService struct{}

func (h *HelloService) mustEmbedUnimplementedHelloServiceServer() {

}

func (h *HelloService) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Resp: []string{req.Req},Timestamp: timestamppb.Now()}, nil
}

func (h *HelloService) SayHelloForServerStream(req *HelloRequest, stream HelloService_SayHelloForServerStreamServer) error {
	for i := 0; i < 10; i++ {
		stream.Send(&HelloResponse{Resp: []string{req.Req}})
		time.Sleep(time.Second)
	}
	return nil
}

func (h *HelloService) SayHelloForClientStream(stream HelloService_SayHelloForClientStreamServer) error {
	resp := make([]string, 10)
	for {
		// 接收客户端发来的流式数据
		res, err := stream.Recv()
		if err == io.EOF {
			// 最终统一回复
			return stream.SendAndClose(&HelloResponse{
				Resp: resp,
			})
		}
		if err != nil {
			return err
		}
		//根据接收到的数据处理
		fmt.Println(res.Req)
		resp = append(resp, res.Req)
	}
	return nil
}

func  (h *HelloService) SayHelloForStream(stream HelloService_SayHelloForStreamServer) error {
	//resp := make([]string, 10)
	for {
		// 接收客户端发来的流式数据
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("客户端发送数据结束")
			//stream.SendAndClose(&HelloResponse{Resp: []string{res.Req}})
			break
		}
		if err != nil {
			fmt.Println(err)//rpc error: code = Canceled desc = context canceled
			return err
		}

		fmt.Println(res.Req)

		err = stream.Send(&HelloResponse{Resp: []string{time.Now().Format("2006-01-02 15:04:05 ") + res.Req}})
		if err != nil {
			return err
		}
	}
	return nil
}
