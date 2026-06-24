package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"backend/pkg/utils"
	"log"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (us *userService) FindAll() {
	us.repo.FindAllUser()
	log.Println("Get all users into user repository")
}
func (us *userService) CreateUser(user model.User) {
	user.Email = utils.NormalizeString(user.Email)
	log.Println("Get all users into user service")
}
func (us *userService) FindByUUID() {
	log.Println("Get all users into user service")
}
func (us *userService) UpdateUser() {
	log.Println("Get all users into user service")
}
func (us *userService) DeleteUser() {
	log.Println("Get all users into user service")
}
