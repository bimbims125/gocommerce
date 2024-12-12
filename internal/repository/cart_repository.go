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

func (r *CartRepository) GetCartByUserID(ctx context.Context, userID int) ([]entity.Cart, error) {
	query := "SELECT id, user_id, product_id, quantity FROM carts WHERE user_id = $1"
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var carts []entity.Cart
	for rows.Next() {
		var cart entity.Cart
		if err := rows.Scan(&cart.ID, &cart.UserID, &cart.ProductID, &cart.Quantity); err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}
