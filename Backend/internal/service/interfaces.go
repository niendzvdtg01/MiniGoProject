package service

import "backend/internal/model"

type UserService interface {
	FindAll(search string, page, limit int) ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	FindByUUID(uuid string) (model.User, error)
	UpdateUser(uuid string, user model.User) (model.User, error)
	DeleteUser()
}
