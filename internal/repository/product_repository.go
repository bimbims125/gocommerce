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

func (r *ProductRepository) GetProducts(ctx context.Context) ([]entity.GetProduct, error) {
	query := `SELECT p.id, p.name, p.price, p.category_id, p.stock, p.sold, p.image_url, c.name as category_name
						FROM products p
						JOIN categories c ON p.category_id = c.id
						ORDER BY p.id ASC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []entity.GetProduct
	for rows.Next() {
		var product entity.GetProduct
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Category.ID, &product.Stock, &product.Sold, &product.ImageURL, &product.Category.Name); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*entity.GetProduct, error) {
	query := `SELECT p.id, p.name, p.price, p.category_id, p.stock, p.sold, p.image_url, c.name as category_name
						FROM products p
						JOIN categories c ON p.category_id = c.id
						WHERE p.id = $1
						ORDER BY p.id ASC
						`
	var product entity.GetProduct
	err := r.db.QueryRowContext(ctx, query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Category.ID, &product.Stock, &product.Sold, &product.ImageURL, &product.Category.Name)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
