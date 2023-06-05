package mapper

import (
	"go-grpc-service/domain"
	"go-grpc-service/shared/proto"
)

func FindAllResponseMapper(in []domain.Villager) *proto.FindAllResponse {

	var villagers []*proto.Villager

	for _, villager := range in {

		protoVillager := &proto.Villager{
			Name:        villager.Name,
			Personality: villager.Personality,
		}
		villagers = append(villagers, protoVillager)
	}

	res := proto.FindAllResponse{
		Villagers: villagers,
	}

	return &res
}
