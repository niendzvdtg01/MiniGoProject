package dto

type PostProduct struct {
	Name string `json:"product_name" binding:"required,min=3,max=100"`
}
