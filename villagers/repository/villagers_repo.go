package repository

import (
	"bufio"
	"context"
	"errors"
	"go-grpc-service/domain"
	"log"
	"os"
	"strings"
)

type VillagersRepository struct {
	Villagers []domain.Villager
}

func NewVillagersRepository() *VillagersRepository {
	var villagerList []domain.Villager
	populateVillagers(&villagerList)

	return &VillagersRepository{Villagers: villagerList}
}

func (v VillagersRepository) FindAll(ctx context.Context) ([]domain.Villager, error) {
	return v.Villagers, nil
}

func (v VillagersRepository) Find(ctx context.Context, name string) (*domain.Villager, error) {

	for _, villager := range v.Villagers {
		if villager.Name == name {
			return &villager, nil
		}
	}
	return nil, errors.New("Not Found")
}

func populateVillagers(list *[]domain.Villager) {

	f, err := os.Open("VillagerList.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data := strings.Fields(scanner.Text())
		villager := domain.Villager{
			Name:        data[0],
			Personality: data[3],
			Species:     data[4],
			Birthday:    data[5] + " " + data[6],
			Catchphrase: data[7],
		}
		*list = append(*list, villager)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
