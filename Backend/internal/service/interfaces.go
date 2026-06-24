package service

import "backend/internal/model"

type UserService interface {
	FindAll()
	CreateUser(user model.User)
	FindByUUID()
	UpdateUser()
	DeleteUser()
}
