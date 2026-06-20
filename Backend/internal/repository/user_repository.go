package repository

import (
	"backend/internal/model"
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

func (ir *InMemoryUserRepository) FindAllUser() {
	log.Println("Get all users into user repo")
}
func (ir *InMemoryUserRepository) CreateUser() {
	log.Println("Get all users into user service")
}
func (ir *InMemoryUserRepository) FindByUUID() {
	log.Println("Get all users into user service")
}
func (ir *InMemoryUserRepository) UpdateUser() {
	log.Println("Get all users into user service")
}
func (ir *InMemoryUserRepository) DeleteUser() {
	log.Println("Get all users into user service")
}
