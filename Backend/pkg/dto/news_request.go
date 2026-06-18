package dto

type NewsFormRequest struct {
	Title  string `form:"title" binding:"required"`
	Status string `form:"status" binding:"required"`
}
