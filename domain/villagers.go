package domain

import (
	"context"
)

type Villager struct {
	Name        string `json:"name"`
	Personality string `json:"personality"`
	Species     string `json:"species"`
	Birthday    string `json:"birthday"`
	Catchphrase string `json:"catchphrase"`
}

type VillagersUsecase interface {
	FindAll(ctx context.Context) ([]Villager, error)
	FindByName(ctx context.Context, name string) (*Villager, error)
}

type VillagersRepository interface {
	FindAll(ctx context.Context) ([]Villager, error)
	FindByName(ctx context.Context, name string) (*Villager, error)
}
