package repository

import (
	"context"
	"database/sql"
	"gocommerce/internal/entity"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) CreateCart(ctx context.Context, cart *entity.Cart) (int, error) {
	query := "INSERT INTO carts (user_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := r.db.QueryRowContext(ctx, query, cart.UserID, cart.ProductID, cart.Quantity).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
