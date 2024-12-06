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
	GetCategoryByID(ctx context.Context, id int) (*entity.Category, error)
}

func NewCategoryUseCase(repo CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		repo: repo,
	}
}
func (u *CategoryUseCase) CreateCategory(ctx context.Context, category *entity.Category) (int, error) {
	if err := category.Validate(); err != nil {
		return 0, err
	}
	return u.repo.CreateCategory(ctx, category)
}

func (u *CategoryUseCase) GetCategories(ctx context.Context) ([]entity.Category, error) {
	return u.repo.GetCategories(ctx)
}

func (u *CategoryUseCase) GetCategoryByID(ctx context.Context, id int) (*entity.Category, error) {
	return u.repo.GetCategoryByID(ctx, id)
}
