package repositories

import (
	"log"

	"github.com/alfanzain/project-sprint-eniqilo-store/src/databases"
	"github.com/alfanzain/project-sprint-eniqilo-store/src/entities"

	"database/sql"
)

type IProductRepository interface {
	Store(*ProductStorePayload) (*entities.Product, error)
}

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository() IProductRepository {
	return &ProductRepository{DB: databases.PostgreSQLInstance}
}

type (
	ProductStorePayload struct {
		Name string
		SKU  string
	}
)

func (r *ProductRepository) Store(p *ProductStorePayload) (*entities.Product, error) {
	var id string
	err := r.DB.QueryRow("INSERT INTO products (name, sku) VALUES ($1, $2) RETURNING id", p.Name, p.SKU).Scan(&id)
	if err != nil {
		log.Printf("Error inserting product: %s", err)
		return nil, err
	}

	product := &entities.Product{
		ID:   id,
		Name: p.Name,
		SKU:  p.SKU,
	}

	return product, nil
}
