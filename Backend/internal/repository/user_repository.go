package repository

import "backend/internal/model"

type InMemoryUserRepository struct {
	user []model.User
}

func NewUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		user: make([]model.User, 0),
	}
}
