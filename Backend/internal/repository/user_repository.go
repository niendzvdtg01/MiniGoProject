package repository

import (
	"backend/internal/model"
	"fmt"
	"log"
)

type InMemoryUserRepository struct {
	user []model.User
}

func NewUserRepository() UserRepository {
	return &InMemoryUserRepository{
		user: make([]model.User, 0),
	}
}

func (ir *InMemoryUserRepository) FindAllUser() ([]model.User, error) {
	return ir.user, nil
}
func (ir *InMemoryUserRepository) CreateUser(user model.User) error {
	ir.user = append(ir.user, user)
	return nil
}
func (ir *InMemoryUserRepository) FindByUUID(uuid string) (model.User, error) {
	for _, user := range ir.user {

		log.Println(len(ir.user))
		if user.UUID == uuid {
			return user, nil
		}
	}
	return model.User{}, nil
}
func (ir *InMemoryUserRepository) UpdateUser(uuid string, users model.User) error {
	for i, user := range ir.user {
		if user.UUID == uuid {
			ir.user[i] = users
			return nil
		}
	}
	return fmt.Errorf("Can not find user in memory!")
}
func (ir *InMemoryUserRepository) DeleteUser() {
	log.Println("Get all users into user service")
}

func (ir *InMemoryUserRepository) FindByEmail(email string) (model.User, bool) {
	for _, user := range ir.user {
		log.Println(email)
		log.Println(ir.user)
		log.Println(len(ir.user))
		if user.Email == email {
			return user, true
		}
	}
	return model.User{}, false
}
