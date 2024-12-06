package repository

import (
	"context"
	"database/sql"
	"gocommerce/internal/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (int, error)
}

type UserRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{db: db}
}

func (r *UserRepositoryPostgres) Create(ctx context.Context, user *entity.User) (int, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepositoryPostgres) GetUsers(ctx context.Context) ([]entity.User, error) {
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
