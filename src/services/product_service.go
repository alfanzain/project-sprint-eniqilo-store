package services

import (
	"errors"

	"github.com/alfanzain/project-sprint-eniqilo-store/src/entities"
	"github.com/alfanzain/project-sprint-eniqilo-store/src/repositories"
)

type IProductService interface {
	AddProduct(*ProductPayload) (*entities.Product, error)
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
		Name        string
		SKU         string
		Category    string
		ImageUrl    string
		Notes       string
		Price       int64
		Stock       int32
		Location    string
		IsAvailable bool
	}
)

func (s *ProductService) AddProduct(p *ProductPayload) (*entities.Product, error) {
	product, err := s.productRepository.Store(&repositories.ProductStorePayload{
		Name:        p.Name,
		SKU:         p.SKU,
		Category:    p.Category,
		ImageUrl:    p.ImageUrl,
		Notes:       p.Notes,
		Price:       p.Price,
		Stock:       p.Stock,
		Location:    p.Location,
		IsAvailable: p.IsAvailable,
	})
	if err != nil {
		return nil, err
	}

	return product, nil
}
