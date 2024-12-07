package repository

import (
	"context"
	"database/sql"
	"gocommerce/internal/entity"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, category *entity.Category) (int, error) {
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id"
	var id int
	err := r.db.QueryRowContext(ctx, query, category.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CategoryRepository) GetCategories(ctx context.Context) ([]entity.Category, error) {
	query := "SELECT id, name FROM categories"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id int) (*entity.Category, error) {
	query := "SELECT id, name FROM categories WHERE id = $1"
	var category entity.Category
	err := r.db.QueryRowContext(ctx, query, id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}
