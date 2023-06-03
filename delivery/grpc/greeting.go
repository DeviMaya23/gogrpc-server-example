package grpc

import (
	"context"
	"go-grpc-service/shared/proto"
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
