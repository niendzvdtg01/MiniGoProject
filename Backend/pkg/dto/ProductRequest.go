package dto

type PostProduct struct {
	Name string `json:"product_name" binding:"required,min=3,max=100"`
}

type ProductRequest struct {
	ProductName string `json:"product_name" binding:"reqired,max=30,min=3"`
	Price       int    `json:"price" binding:"required"`
	BestSaling  bool   `json:"best_saling" binding:"required"`
}
