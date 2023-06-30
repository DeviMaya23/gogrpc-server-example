package usecase

import (
	"context"
	"go-grpc-service/domain"
)

type VillagersUsecase struct {
	repo domain.VillagersRepository
}

func NewVillagersUsecase(repo domain.VillagersRepository) *VillagersUsecase {
	return &VillagersUsecase{
		repo: repo,
	}
}

func (v VillagersUsecase) FindAll(ctx context.Context) ([]domain.Villager, error) {
	res, err := v.repo.FindAll(ctx)
	return res, err
}

func (v VillagersUsecase) Find(ctx context.Context, name string) (*domain.Villager, error) {
	res, err := v.repo.Find(ctx, name)
	return res, err
}
