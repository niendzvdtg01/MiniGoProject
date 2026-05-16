package dto

type PostProduct struct {
	Name         string       `json:"product_name" binding:"required,min=3,max=100"`
	ProductImage ProductImage `json:"product_image" binding:"required"`
	Display      bool         `json:"display"`
}

type ProductRequest struct {
	ProductName string `json:"product_name" binding:"required,max=30,min=3"`
	Price       int    `json:"price" binding:"required,min_int=1,max_int=1000000"`
	BestSaling  bool   `json:"best_saling" binding:"required"`
}

type ProductImage struct {
	ImageName string `json:"image_name" binding:"required"`
	ImageLink string `json:"image_link" binding:"required"`
}
