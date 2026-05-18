package service

import (
	"Backend/pkg/dto"

	"github.com/google/uuid"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

// lol
func (ps *ProductService) ValidateUUID(product dto.PostProduct) error {
	for key := range product.ProductInfo {
		if _, err := uuid.Parse(key); err != nil {
			return err
		}
	}
	return nil
}
