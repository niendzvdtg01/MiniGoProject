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
