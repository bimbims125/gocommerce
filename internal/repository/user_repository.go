package repository

import (
	"context"
	"database/sql"
	"gocommerce/internal/entity"
)

// type UserRepository interface {
// 	Create(ctx context.Context, user *entity.User) (int, error)
// 	GetUsers(ctx context.Context) ([]entity.User, error)
// 	GetUserByID(ctx context.Context, id int) (*entity.User, error)
// }

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (int, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]entity.User, error) {
	query := "SELECT id, name, email, password FROM users"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE id = $1"
	var user entity.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
