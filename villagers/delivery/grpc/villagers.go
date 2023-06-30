package grpc

import (
	"context"
	"go-grpc-service/domain"
	"go-grpc-service/shared/proto"
	"go-grpc-service/villagers/mapper"
	"io"
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

// FindAllStreamServerSide(*emptypb.Empty, "go-grpc-service/shared/proto".VillagersService_FindAllStreamServerSideServer) error
func (v VillagersHandler) FindAllStreamServerSide(in *empty.Empty, stream proto.VillagersService_FindAllStreamServerSideServer) error {

	ctx := context.Background()
	res, err := v.villagersUsecase.FindAll(ctx)
	if err != nil {
		return err
	}

	for _, villager := range res {
		stream.Send(mapper.FindAllStreamServerSideMapper(villager))
		time.Sleep(50 * time.Millisecond)
	}

	return nil
}

func (v VillagersHandler) FindStreamClientSide(stream proto.VillagersService_FindStreamClientSideServer) error {

	ctx := context.Background()

	var result []domain.Villager
	for {
		name, err := stream.Recv()

		// read stream
		if err == io.EOF {
			res := mapper.FindAllResponseMapper(result)
			return stream.SendAndClose(&proto.FindAllResponse{
				Villagers: res.Villagers,
			})
		}
		if err != nil {
			return err
		}

		searchResult, _ := v.villagersUsecase.Find(ctx, name.Name)
		if searchResult != nil {
			result = append(result, *searchResult)
		}
	}
}
