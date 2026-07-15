package dto

import (
	"backend/internal/model"
)

type UserRequest struct {
	UUID     string `json:"uuid"`
	FullName string `json:"name" `
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Status   string `json:"status"`
	Level    string `json:"level"`
}

type CreateUserInput struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,email_advance"`
	Age      int    `json:"age" binding:"required,gt=0"`
	Password string `json:"password" binding:"required,min=8,password_strong"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}
type UpdateUserInput struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty,email,email_advance"`
	Age      int    `json:"age" binding:"omitempty,gt=0"`
	Password string `json:"password" binding:"omitempty,min=8,password_strong"`
	Status   int    `json:"status" binding:"omitempty,oneof=1 2"`
	Level    int    `json:"level" binding:"omitempty,oneof=1 2"`
}

func MapUserToDTO(user model.User) *UserRequest {
	return &UserRequest{
		UUID:     user.UUID,
		FullName: user.Name,
		Email:    user.Email,
		Age:      user.Age,
		Status:   mapStatusTest(user.Status),
		Level:    mapLevelTest(user.Level),
	}
}

func (c *CreateUserInput) ToModel() model.User {
	return model.User{}
}

func MapUserToDTOs(users []model.User) []UserRequest {
	dtos := make([]UserRequest, len(users))
	for _, user := range users {
		dto := UserRequest{
			UUID:     user.UUID,
			FullName: user.Name,
			Email:    user.Email,
			Age:      user.Age,
			Status:   mapStatusTest(user.Status),
			Level:    mapLevelTest(user.Level),
		}
		dtos = append(dtos, dto)
	}
	return dtos
}

func mapStatusTest(status int) string {
	switch status {
	case 1:
		return "Show"
	case 2:
		return "Hide"
	default:
		return "None"
	}
}

func mapLevelTest(level int) string {
	switch level {
	case 1:
		return "Admin"
	case 2:
		return "Member"
	default:
		return "None"
	}
}
