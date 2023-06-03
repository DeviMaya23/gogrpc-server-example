package grpc

import (
	"context"
	"go-grpc-service/shared/proto/greeting"
)

type GreetingHandler struct {
	greeting.UnimplementedGreetingServiceServer
}

func NewGreetingHandler() *GreetingHandler {
	return &GreetingHandler{}
}

func (g GreetingHandler) GetNamedGreeting(ctx context.Context, req *greeting.GetNamedGreetingRequest) (*greeting.GreetingResponse, error) {

	message := "Hello " + req.Name
	return &greeting.GreetingResponse{
		Message: message,
	}, nil
}
