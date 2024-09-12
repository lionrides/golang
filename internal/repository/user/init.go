package user

import (
	"context"
	"database/sql"
	"golang/internal/domain"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepo{
		DB: db,
	}
}

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]domain.User, error)
	CreateUser(domain.User) (string, error)
	UpdateUser(id int, user domain.User) (string, error)
	DeleteUser(int) (string, error)
}
