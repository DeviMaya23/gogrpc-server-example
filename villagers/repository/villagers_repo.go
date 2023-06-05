package repository

import (
	"context"
	"go-grpc-service/domain"
)

type VillagersRepository struct {
	Villagers []domain.Villager
}

func NewVillagersRepository() *VillagersRepository {
	var villagerList []domain.Villager
	villagerList = append(villagerList, domain.Villager{
		Name:        "Merengue",
		Personality: "Normal",
	})

	villagerList = append(villagerList, domain.Villager{
		Name:        "Chadder",
		Personality: "Smug",
	})

	villagerList = append(villagerList, domain.Villager{
		Name:        "Billy",
		Personality: "Jock",
	})

	villagerList = append(villagerList, domain.Villager{
		Name:        "Twiggy",
		Personality: "Peppy",
	})

	return &VillagersRepository{Villagers: villagerList}
}

func (v VillagersRepository) FindAll(ctx context.Context) ([]domain.Villager, error) {
	return v.Villagers, nil
}
