package usecase

import (
	"context"
	"gocommerce/internal/entity"
)

type OrderUsecase struct {
	repo OrderRepository
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *entity.Order) (int, error)
	GetOrders(ctx context.Context) ([]entity.GetOrder, error)
	UpdateOrderPaymentStatus(ctx context.Context, transactionID string, paymentStatus string) error
}

func NewOrderUsecase(repo OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: repo}
}

func (u *OrderUsecase) CreateOrder(ctx context.Context, order *entity.Order) (int, error) {
	return u.repo.CreateOrder(ctx, order)
}

func (u *OrderUsecase) GetOrders(ctx context.Context) ([]entity.GetOrder, error) {
	return u.repo.GetOrders(ctx)
}

func (u *OrderUsecase) UpdateOrderPaymentStatus(ctx context.Context, transactionID string, paymentStatus string) error {
	return u.repo.UpdateOrderPaymentStatus(ctx, transactionID, paymentStatus)
}
