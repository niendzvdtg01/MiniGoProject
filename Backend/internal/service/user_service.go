package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"backend/pkg/utils"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
func (us *userService) CreateUser(user model.User) (model.User, error) {
	user.Email = utils.NormalizeString(user.Email)
	if _, exists := us.repo.FindByEmail(user.Email); exists {
		return model.User{}, utils.NewError("email already exists", utils.ErrCodeConflict)
	}

	user.UUID = uuid.New().String()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, utils.WrapError("failded to create user", utils.ErrCodeInternal, err)
	}
	//
	user.Password = string(passwordHash)

	us.repo.CreateUser(user)
	return user, nil
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
