package repositories

import (
	"github.com/alfanzain/project-sprint-eniqilo-store/src/databases"
	"github.com/alfanzain/project-sprint-eniqilo-store/src/entities"

	"database/sql"
	"errors"
	"log"
)

type IStaffRepository interface {
	FindByPhoneNumber(string) (*entities.Staff, error)
	DoesPhoneExist(string) (bool, error)
	Store(*StaffStorePayload) (*entities.Staff, error)
}

type StaffRepository struct {
	DB *sql.DB
}

func NewStaffRepository() IStaffRepository {
	return &StaffRepository{DB: databases.PostgreSQLInstance}
}

type (
	StaffStorePayload struct {
		ID          string
		Name        string
		PhoneNumber string
		Password    string
	}
)

func (r *StaffRepository) FindByPhoneNumber(phoneNumber string) (*entities.Staff, error) {
	var staff entities.Staff
	err := r.DB.QueryRow(`SELECT id, phone_number, name, password FROM users WHERE phone_number = $1`, phoneNumber).Scan(&staff.ID, &staff.PhoneNumber, &staff.Name, &staff.Password)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalln(err)
		return nil, err
	}

	return &staff, nil
}

func (r *StaffRepository) DoesPhoneExist(phoneNumber string) (bool, error) {
	var scannedPhoneNumber string
	err := r.DB.QueryRow(`SELECT phone_number FROM users WHERE phone_number = $1`, phoneNumber).Scan(&scannedPhoneNumber)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalln(err)
		return false, err
	}

	if len(scannedPhoneNumber) == 0 {
		return false, nil
	}

	return true, nil
}

func (r *StaffRepository) Store(p *StaffStorePayload) (*entities.Staff, error) {
	var id string
	err := r.DB.QueryRow("INSERT INTO users (name, phone_number, password) VALUES ($1, $2, $3) RETURNING id", p.Name, p.PhoneNumber, p.Password).Scan(&id)
	if err != nil {
		log.Printf("Error inserting staff: %s", err)
		return nil, err
	}

	staff := &entities.Staff{
		ID:          id,
		Name:        p.Name,
		PhoneNumber: p.PhoneNumber,
	}

	return staff, nil
}
