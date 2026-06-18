package service

import (
	"backend/pkg/dto"

	"github.com/google/uuid"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) ValidateProductInfoKeys(product dto.CreateProductRequest) error {
	for key := range product.ProductInfo {
		if _, err := uuid.Parse(key); err != nil {
			return err
		}
	}
	return nil
}
