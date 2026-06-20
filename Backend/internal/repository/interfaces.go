package repository

type UserRepository interface {
	FindAllUser()
	CreateUser()
	FindByUUID()
	UpdateUser()
	DeleteUser()
}
