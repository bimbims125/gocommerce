package repository

import (
	"context"
	"database/sql"
	"gocommerce/internal/entity"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *entity.Product) (int, error) {
	query := "INSERT INTO products (name,  price, category_id, stock, image_url) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := r.db.QueryRowContext(ctx, query, product.Name, product.Price, product.CategoryID, product.Stock, product.ImageURL).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
