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
	GetProducts(ctx context.Context) ([]entity.GetProduct, error)
	GetProductByID(ctx context.Context, id int) (*entity.GetProduct, error)
}

func NewProductUsecase(repo ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (u *ProductUseCase) CreateProduct(ctx context.Context, product *entity.Product) (int, error) {
	return u.repo.CreateProduct(ctx, product)
}

func (u *ProductUseCase) GetProducts(ctx context.Context) ([]entity.GetProduct, error) {
	return u.repo.GetProducts(ctx)
}

func (u *ProductUseCase) GetProductByID(ctx context.Context, id int) (*entity.GetProduct, error) {
	return u.repo.GetProductByID(ctx, id)
}
