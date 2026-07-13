package repository

import (
	"backend/internal/model"
)

type UserRepository interface {
	FindAllUser() ([]model.User, error)
	CreateUser(user model.User) error
	FindByUUID(uuid string) (model.User, error)
	UpdateUser(uuid string, user model.User) error
	DeleteUser()
	FindByEmail(email string) (model.User, bool)
}
