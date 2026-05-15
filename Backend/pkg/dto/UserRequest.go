package dto

type UserRequest struct {
	UserName string `json:"username" binding:"required,max=100,min=3"`
	Password string `json:"password" binding:"required,max=30,min=3"`
}


