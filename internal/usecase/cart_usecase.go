package usecase

import (
	"context"
	"gocommerce/internal/entity"
	"gocommerce/internal/repository"
)

type CartUsecase struct {
	repo repository.CartRepository
}

type CartRepository interface {
	CreateCart(ctx context.Context, cart *entity.Cart) (int, error)
}

func NewCartUsecase(repo repository.CartRepository) *CartUsecase {
	return &CartUsecase{repo: repo}
}

func (u *CartUsecase) CreateCart(ctx context.Context, cart *entity.Cart) (int, error) {
	return u.repo.CreateCart(ctx, cart)
}