package repository

import "backend/internal/model"

type UserRepository interface {
	FindAllUser() ([]model.User, error)
	CreateUser(user model.User)
	FindByUUID(uuid string) (model.User, error)
	UpdateUser()
	DeleteUser()
	FindByEmail(email string) (model.User, bool)
}
