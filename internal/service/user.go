package service

import (
	"context"
	"database/sql"
	"golang/internal/domain"
	"golang/internal/repository/user"
)

type UserService struct {
	userRepo user.UserRepository
}

func NewUserService(userRepo user.UserRepository) UserService {
	return &UserService(userRepo:userRepo)
}

func GetAllUser