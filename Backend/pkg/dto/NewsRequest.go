package dto

type PostNewsV1 struct {
	Title  string `form:"title" binding:"required"`
	Status string `form:"status" binding:"required"`
}
