package usecase

import (
	"context"
	"gocommerce/internal/entity"
)

type CategoryUseCase struct {
	repo CategoryRepository
}

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *entity.Category) (int, error)
	GetCategories(ctx context.Context) ([]entity.Category, error)
}

func NewCategoryUseCase(repo CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		repo: repo,
	}
}
func (c *CategoryUseCase) CreateCategory(ctx context.Context, category *entity.Category) (int, error) {
	if err := category.Validate(); err != nil {
		return 0, err
	}
	return c.repo.CreateCategory(ctx, category)
}

func (c *CategoryUseCase) GetCategories(ctx context.Context) ([]entity.Category, error) {
	return c.repo.GetCategories(ctx)
}
