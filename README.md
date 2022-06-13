
> go get -u google.golang.org/grpc


protoc 生成文件
> protoc --go_out=./ --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false proto/Hello.proto