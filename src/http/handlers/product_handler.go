package handlers

import (
	"net/http"

	"github.com/alfanzain/project-sprint-eniqilo-store/src/repositories"
	"github.com/alfanzain/project-sprint-eniqilo-store/src/services"
	"github.com/labstack/echo/v4"
)

type IProductHandler interface {
	Store(c echo.Context) error
}

type ProductHandler struct {
	productService services.IProductService
}

func NewProductHandler(s services.IProductService) IProductHandler {
	return &ProductHandler{
		productService: services.NewProductService(
			repositories.NewProductRepository(),
		),
	}
}

type (
	ProductRequest struct {
		Name        string `json:"name" validate:"required,min=1,max=30"`
		SKU         string `json:"sku" validate:"required,min=1,max=30"`
		Category    string `json:"category" validate:"required"` // validate=enum?
		ImageUrl    string `json:"imageUrl" validate:"required"`
		Notes       string `json:"notes" validate:"required,min=1,max=200"`
		Price       int64  `json:"price" validate:"required,min=1"`
		Stock       int32  `json:"stock" validate:"required,min=0,max=100000"`
		Location    string `json:"location" validate:"required,min=1,max=200"`
		IsAvailable bool   `json:"isAvailable" validate:"required"`
	}
)

func (h *ProductHandler) Store(c echo.Context) (e error) {
	r := new(ProductRequest)

	if e = c.Bind(r); e != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: e.Error(),
		})
	}

	if e = c.Validate(r); e != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: e.Error(),
		})
	}

	data, err := h.productService.AddProduct(&services.ProductPayload{
		Name:        r.Name,
		SKU:         r.SKU,
		Category:    r.Category,
		ImageUrl:    r.ImageUrl,
		Notes:       r.Notes,
		Price:       r.Price,
		Stock:       r.Stock,
		Location:    r.Location,
		IsAvailable: r.IsAvailable,
	})

	if err != nil {
		return c.JSON(http.StatusConflict, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "success",
		Data:    data,
	})
}
