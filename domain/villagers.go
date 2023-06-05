package domain

import (
	"context"
)

type Villager struct {
	Name        string `json:"id"`
	Personality string `json:"title"`
}

type VillagersUsecase interface {
	FindAll(ctx context.Context) ([]Villager, error)
}

type VillagersRepository interface {
	FindAll(ctx context.Context) ([]Villager, error)
}
