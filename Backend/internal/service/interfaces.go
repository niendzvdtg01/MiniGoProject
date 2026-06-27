package service

import "backend/internal/model"

type UserService interface {
	FindAll()
	CreateUser(user model.User) (model.User, error)
	FindByUUID()
	UpdateUser()
	DeleteUser()
}
