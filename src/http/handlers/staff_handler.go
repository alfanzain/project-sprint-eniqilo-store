package handlers

import (
	"net/http"

	"github.com/alfanzain/eniqilo-store/src/repositories"
	"github.com/alfanzain/eniqilo-store/src/services"

	"github.com/labstack/echo/v4"
)

type IStaffHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type StaffHandler struct {
	staffService services.IStaffService
}

func NewStaffHandler(s services.IStaffService) IStaffHandler {
	return &StaffHandler{
		staffService: services.NewStaffService(
			repositories.NewStaffRepository(),
		),
	}
}

type (
	RegisterRequest struct {
		PhoneNumber string `json:"phoneNumber" validate:"required"`
		Name        string `json:"name" validate:"required,min=5,max=50"`
		Password    string `json:"password" validate:"required,min=5,max=15"`
	}

	LoginRequest struct {
		PhoneNumber string `json:"phoneNumber" validate:"required"`
		Password    string `json:"password" validate:"required,min=5,max=15"`
	}
)

func (h *StaffHandler) Register(c echo.Context) (e error) {
	r := new(RegisterRequest)

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

	data, err := h.staffService.Register(&services.RegisterPayload{
		PhoneNumber: r.PhoneNumber,
		Name:        r.Name,
		Password:    r.Password,
	})

	if err != nil {
		return c.JSON(http.StatusConflict, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User registered successfully",
		Data:    data,
	})
}

func (h *StaffHandler) Login(c echo.Context) (e error) {
	r := new(LoginRequest)

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

	data, e := h.staffService.Login(&services.LoginPayload{
		PhoneNumber: r.PhoneNumber,
		Password:    r.Password,
	})

	if e != nil {
		if e == services.ErrStaffNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: e.Error(),
			})
		} else if e == services.ErrInvalidPassword {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: e.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "User logged successfully",
		Data:    data,
	})
}
