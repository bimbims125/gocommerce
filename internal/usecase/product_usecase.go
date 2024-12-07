package usecase

import (
	"context"
	"gocommerce/internal/entity"
)

type ProductUseCase struct {
	repo ProductRepository
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *entity.Product) (int, error)
}

func NewProductUsecase(repo ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (u *ProductUseCase) CreateProduct(ctx context.Context, product *entity.Product) (int, error) {
	return u.repo.CreateProduct(ctx, product)
}
