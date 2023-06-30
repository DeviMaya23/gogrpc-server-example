package grpc

import (
	"context"
	"go-grpc-service/common/constants"
	"go-grpc-service/domain"
	"go-grpc-service/shared/proto"
	"go-grpc-service/villagers/mapper"
	"io"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (v VillagersHandler) FindByName(ctx context.Context, req *proto.FindByNameRequest) (*proto.Villager, error) {

	res, err := v.villagersUsecase.FindByName(ctx, req.Name)
	if err != nil {
		switch {
		case err.Error() == constants.ErrorNotFound:
			return nil, status.Errorf(codes.NotFound, constants.ErrorNotFound)
		default:
			return nil, status.Errorf(codes.Internal, constants.ErrorGeneric)
		}
	}

	return mapper.VillagerMapper(res), nil
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

		searchResult, _ := v.villagersUsecase.FindByName(ctx, name.Name)
		if searchResult != nil {
			result = append(result, *searchResult)
		}
	}
}

func (v VillagersHandler) FindStreamBidirecitonal(stream proto.VillagersService_FindStreamBidirecitonalServer) error {

	ctx := stream.Context()

	for {
		in, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		searchResult, errSearch := v.villagersUsecase.FindByName(ctx, in.Name)
		if errSearch == nil {
			if err := stream.Send(mapper.VillagerMapper(searchResult)); err != nil {
				return err
			}
		}
	}

}
