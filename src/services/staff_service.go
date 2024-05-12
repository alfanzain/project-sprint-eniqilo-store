package services

import (
	"errors"
	"os"

	"github.com/alfanzain/project-sprint-eniqilo-store/src/helpers"
	"github.com/alfanzain/project-sprint-eniqilo-store/src/repositories"
)

type IStaffService interface {
	Register(*RegisterPayload) (*RegisterResult, error)
	Login(*LoginPayload) (*LoginResult, error)
}

type StaffService struct {
	staffRepository repositories.IStaffRepository
}

func NewStaffService(
	staffRepository repositories.IStaffRepository,
) IStaffService {
	return &StaffService{
		staffRepository: staffRepository,
	}
}

var ErrPhoneAlreadyUsed = errors.New("phone number exists")
var ErrStaffNotFound = errors.New("staff not found")
var ErrInvalidPassword = errors.New("invalid password")

type (
	RegisterPayload struct {
		PhoneNumber string
		Name        string
		Password    string
	}

	RegisterResult struct {
		UserID      string `json:"userId"`
		PhoneNumber string `json:"phoneNumber"`
		Name        string `json:"name"`
		Password    string `json:"-"`
		AccessToken string `json:"accessToken"`
	}

	LoginPayload struct {
		PhoneNumber string
		Name        string
		Password    string
	}

	LoginResult struct {
		UserID      string `json:"userId"`
		PhoneNumber string `json:"phoneNumber"`
		Name        string `json:"name"`
		AccessToken string `json:"accessToken"`
	}
)

func (s *StaffService) Register(p *RegisterPayload) (*RegisterResult, error) {
	phoneExists, _ := s.staffRepository.DoesPhoneExist(p.PhoneNumber)

	if phoneExists {
		return nil, ErrPhoneAlreadyUsed
	}

	hashedPassword, _ := helpers.HashPassword(p.Password)
	staff, err := s.staffRepository.Store(&repositories.StaffStorePayload{
		PhoneNumber: p.PhoneNumber,
		Name:        p.Name,
		Password:    hashedPassword,
	})

	paramsGenerateJWTRegister := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 480,
		SecretKey:       os.Getenv("JWT_SECRET"),
		UserId:          staff.ID,
	}

	accessToken, errAccessToken := helpers.GenerateJWT(&paramsGenerateJWTRegister)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	if err != nil {
		return nil, err
	}

	return &RegisterResult{
		UserID:      staff.ID,
		PhoneNumber: p.PhoneNumber,
		Name:        p.Name,
		AccessToken: accessToken,
	}, nil
}

func (s *StaffService) Login(p *LoginPayload) (*LoginResult, error) {
	staff, err := s.staffRepository.FindByPhoneNumber(p.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if staff == nil {
		return nil, ErrStaffNotFound
	}

	isValidPassword := helpers.CheckPasswordHash(p.Password, staff.Password)
	if !isValidPassword {
		return nil, ErrInvalidPassword
	}

	paramsGenerateJWTLogin := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 480,
		SecretKey:       os.Getenv("JWT_SECRET"),
		UserId:          staff.ID,
	}

	accessToken, err := helpers.GenerateJWT(&paramsGenerateJWTLogin)
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		UserID:      staff.ID,
		PhoneNumber: p.PhoneNumber,
		Name:        staff.Name,
		AccessToken: accessToken,
	}, nil
}
