package service

import (
	"Backend/internal/model"

	"github.com/gin-gonic/gin"
)

type UserRepo interface {
	FindUserByID(ctx *gin.Context) (*model.User, error)
	FindUserByEmail(ctx *gin.Context, email string) (*model.User, error)
	CreateUser(ctx *gin.Context, user *model.User) error
}

type AuthService struct {
	userRepo UserRepo
}

func NewAuthService(repo UserRepo) *AuthService {
	return &AuthService{
		userRepo: repo,
	}
}

func (s *AuthService) Login(ctx *gin.Context, email string, password string) (*model.User, error) {
	user, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Register(ctx *gin.Context, user *model.User) error {
	return s.userRepo.CreateUser(ctx, user)
}
