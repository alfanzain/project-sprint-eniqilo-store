package services

import (
	"errors"

	"github.com/alfanzain/eniqilo-store/src/repositories"
)

type IProductService interface {
	AddProduct(*ProductPayload) (*ProductResult, error)
}

type ProductService struct {
	productRepository repositories.IProductRepository
}

func NewProductService(
	productRepository repositories.IProductRepository,
) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

var ErrProductNotFound = errors.New("product not found")

type (
	ProductPayload struct {
		Name string
		SKU  string
	}

	ProductResult struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		SKU  string `json:"sku"`
	}
)

func (s *ProductService) AddProduct(p *ProductPayload) (*ProductResult, error) {
	product, err := s.productRepository.Store(&repositories.ProductStorePayload{
		Name: p.Name,
		SKU:  p.SKU,
	})
	if err != nil {
		return nil, err
	}

	return &ProductResult{
		ID:   product.ID,
		Name: p.Name,
		SKU:  p.SKU,
	}, nil
}
