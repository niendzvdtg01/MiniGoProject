package dto

type PostProduct struct {
	Name             string                 `json:"product_name" binding:"required,min=3,max=100"`
	ProductImage     ProductImage           `json:"product_image" binding:"required"`
	Display          bool                   `json:"display"`
	Tags             []string               `json:"tag" binding:"required,gt=0,lt=5"`
	ProductAttribute []ProductAttribute     `json:"product_attribute" binding:"required,gt=0,dive"`
	ProductInfo      map[string]ProductInfo `json:"product_info" binding:"required"`
	ProductMetaData  map[string]any         `json:"product_metadata" binding:"omitempty"`
}

type ProductRequest struct {
	ProductName string `json:"product_name" binding:"required,max=30,min=3"`
	Price       int    `json:"price" binding:"required,min_int=1,max_int=1000000"`
	BestSaling  bool   `json:"best_saling" binding:"required"`
}

type ProductImage struct {
	ImageName string `json:"image_name" binding:"required,file_extension=jpg png mp4"`
	ImageLink string `json:"image_link" binding:"required"`
}

type ProductAttribute struct {
	AttributeName  string `json:"attribute_name" binding:"required"`
	AttributeValue string `json:"attribute_value" binding:"required"`
}

type ProductInfo struct {
	InfoKey   string `json:"info_key" binding:"required"`
	InfoValue string `json:"info_value" binding:"required"`
}
