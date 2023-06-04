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

	message := fmt.Sprintf("My name is %s! I am %d year(s) old. I like playing ", req.Name, req.Age)

	for _, game := range req.FavoriteGames {
		message = message + fmt.Sprintf("%s(%s), ", game.Name, game.Console)
	}

	message = message + " amongst other things."
	isOld := req.Age > 18

	return &proto.GetVerboseGreetingResponse{
		Message: message,
		IsOld:   isOld,
	}, nil
}
