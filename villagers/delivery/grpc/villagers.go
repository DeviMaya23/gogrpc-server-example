package grpc

import (
	"context"
	"go-grpc-service/domain"
	"go-grpc-service/shared/proto"
	"go-grpc-service/villagers/mapper"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
)

type VillagersHandler struct {
	proto.UnimplementedVillagersServiceServer
	villagersUsecase domain.VillagersUsecase
}

func NewVillagersHandler(uc domain.VillagersUsecase) *VillagersHandler {
	return &VillagersHandler{villagersUsecase: uc}
}

func (v VillagersHandler) FindAll(ctx context.Context, req *empty.Empty) (*proto.FindAllResponse, error) {

	res, err := v.villagersUsecase.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FindAllResponseMapper(res), nil
}

// FindAllStream(*emptypb.Empty, "go-grpc-service/shared/proto".VillagersService_FindAllStreamServer) error
func (v VillagersHandler) FindAllStream(in *empty.Empty, stream proto.VillagersService_FindAllStreamServer) error {

	ctx := context.Background()
	res, err := v.villagersUsecase.FindAll(ctx)
	if err != nil {
		return err
	}

	for _, villager := range res {
		stream.Send(mapper.FindAllStreamMapper(villager))
		time.Sleep(50 * time.Millisecond)
	}

	return nil
}
