package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"backend/pkg/utils"
	"log"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (us *userService) FindAll(search string, page, limit int) ([]model.User, error) {
	user, err := us.repo.FindAllUser()
	if err != nil {
		return nil, utils.WrapError("failed to fetch user", utils.ErrCodeInternal, err)
	}

	var filterUsers []model.User

	if search != "" {
		search = strings.ToLower(search)
		for _, u := range user {
			name := strings.ToLower(u.Name)
			email := strings.ToLower(u.Email)

			if strings.Contains(name, search) || strings.Contains(email, search) {
				filterUsers = append(filterUsers, u)
			}

		}
	} else {
		filterUsers = user
	}
	start := (page - 1) * limit
	if start >= len(filterUsers) {
		return []model.User{}, nil
	}

	end := start + limit

	if end > len(filterUsers) {
		end = len(filterUsers)
	}

	log.Println("Get all users into user repository")
	return filterUsers[start:end], nil
}
func (us *userService) CreateUser(user model.User) (model.User, error) {
	user.Email = utils.NormalizeString(user.Email)
	if _, exists := us.repo.FindByEmail(user.Email); exists {
		log.Println(exists)
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
func (us *userService) FindByUUID(uuid string) (model.User, error) {
	user, err := us.repo.FindByUUID(uuid)
	if err != nil {
		return model.User{}, utils.NewError("user not found", utils.ErrCodeBadRequest)
	}
	return user, nil
}
func (us *userService) UpdateUser(uuid string, user model.User) (model.User, error) {
	user.Email = utils.NormalizeString(user.Email)
	if _, exists := us.repo.FindByEmail(user.Email); exists {
		log.Println(exists)
		return model.User{}, utils.NewError("email already exists", utils.ErrCodeConflict)
	}

	currentUser, err := us.repo.FindByUUID(uuid)

	if err != nil {
		return model.User{}, utils.NewError("user not found!!", utils.ErrCodeNotFound)
	}
	currentUser.Name = user.Name
	currentUser.Age = user.Age
	currentUser.Email = user.Email
	currentUser.Level = user.Level
	currentUser.Status = user.Status
	// currentUser.Password = user.Password error

	if currentUser.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return model.User{}, utils.WrapError("Can not hash password!!!", utils.ErrCodeInternal, err)
		}
		currentUser.Password = string(passwordHash)
	}

	if err := us.repo.UpdateUser(uuid, user); err != nil {
		return model.User{}, utils.NewError("Can not update user", utils.ErrCodeInternal)
	}
	return currentUser, nil
}
func (us *userService) DeleteUser() {
	log.Println("Get all users into user service")
}
