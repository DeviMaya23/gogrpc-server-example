package grpc

import (
	"context"
	"fmt"
	"go-grpc-service/shared/proto"

	"github.com/golang/protobuf/ptypes/empty"
)

type GreetingHandler struct {
	proto.UnimplementedGreetingServiceServer
}

func NewGreetingHandler() *GreetingHandler {
	return &GreetingHandler{}
}

func (g GreetingHandler) GetNamedGreeting(ctx context.Context, req *proto.GetNamedGreetingRequest) (*proto.GreetingResponse, error) {

	message := "Hello " + req.Name
	return &proto.GreetingResponse{
		Message: message,
	}, nil
}

func (g GreetingHandler) GetGenericGreeting(ctx context.Context, in *empty.Empty) (*proto.GreetingResponse, error) {

	message := "Hello to you!"
	return &proto.GreetingResponse{
		Message: message,
	}, nil
}

func (g GreetingHandler) GetVerboseGreeting(ctx context.Context, req *proto.GetVerboseGreetingRequest) (*proto.GetVerboseGreetingResponse, error) {

	message := fmt.Sprintf("My name is %s! I am %d year(s) old. My favorite game is %s on the %s. ", req.Name, req.Age, req.FavoriteGame.Name, req.FavoriteGame.Console)
	isOld := req.Age > 18

	return &proto.GetVerboseGreetingResponse{
		Message: message,
		IsOld:   isOld,
	}, nil
}
