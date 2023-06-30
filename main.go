package main

import (
	"flag"
	"fmt"
	greetingDelivery "go-grpc-service/greeting/delivery/grpc"
	"go-grpc-service/shared/proto"
	villagersDelivery "go-grpc-service/villagers/delivery/grpc"
	"go-grpc-service/villagers/repository"
	"go-grpc-service/villagers/usecase"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	villagersRepo := repository.NewVillagersRepository()
	villagersUc := usecase.NewVillagersUsecase(villagersRepo)

	server := grpc.NewServer()
	proto.RegisterVillagersServiceServer(server, villagersDelivery.NewVillagersHandler(villagersUc))
	proto.RegisterGreetingServiceServer(server, greetingDelivery.NewGreetingHandler())

	// Enable Reflection
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
