# Go GRPC Server

## Install Protobuf Compiler
1. Download the latest Protobuf Compiler binary at :

https://github.com/protocolbuffers/protobuf/releases

2. Update `PATH` environment variable to include the path to the `protoc` executable (folder ./bin).

## Install Go plugins for protobuf
```
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
