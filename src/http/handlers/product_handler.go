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
		Name string `json:"name" validate:"required,min=1,max=30"`
		SKU  string `json:"sku" validate:"required,min=1,max=30"`
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
		Name: r.Name,
		SKU:  r.SKU,
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
