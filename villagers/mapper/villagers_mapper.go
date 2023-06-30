package mapper

import (
	"go-grpc-service/domain"
	"go-grpc-service/shared/proto"
)

func VillagerMapper(in *domain.Villager) *proto.Villager {

	protoVillager := &proto.Villager{
		Name:        in.Name,
		Personality: in.Personality,
	}

	return protoVillager
}

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

func FindAllStreamServerSideMapper(in domain.Villager) *proto.Villager {

	protoVillager := &proto.Villager{
		Name:        in.Name,
		Personality: in.Personality,
	}

	return protoVillager
}
