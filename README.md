
> go get -u google.golang.org/grpc


protoc ็ๆๆไปถ
> protoc --go_out=./ --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false proto/Hello.proto