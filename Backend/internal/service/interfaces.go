package service

import "backend/internal/model"

type UserService interface {
	FindAll() ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	FindByUUID()
	UpdateUser()
	DeleteUser()
}
