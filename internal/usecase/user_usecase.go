package usecase

import (
	"context"
	"gocommerce/internal/entity"
	"log"
)

type UserUseCase struct {
	repo UserRepository
}

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (int, error)
	GetUsers(ctx context.Context) ([]entity.User, error)
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) Create(ctx context.Context, user *entity.User) (int, error) {
	if err := user.Validate(); err != nil {
		return 0, err
	}
	// Hash password
	if err := user.HashPassword(); err != nil {
		log.Println(err)
		return 0, err
	}
	return u.repo.Create(ctx, user)
}

func (u *UserUseCase) GetUsers(ctx context.Context) ([]entity.User, error) {
	return u.repo.GetUsers(ctx)
}

func (u *UserUseCase) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	return u.repo.GetUserByID(ctx, id)
}

func (u *UserUseCase) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return u.repo.GetUserByEmail(ctx, email)
}
