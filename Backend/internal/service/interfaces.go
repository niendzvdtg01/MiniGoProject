package service

type UserService interface {
	FindAll()
	CreateUser()
	FindByUUID()
	UpdateUser()
	DeleteUser()
}
