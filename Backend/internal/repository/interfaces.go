package repository

import "backend/internal/model"

type UserRepository interface {
	FindAllUser()
	CreateUser(user model.User)
	FindByUUID()
	UpdateUser()
	DeleteUser()
	FindByEmail(email string) (model.User, bool)
}
