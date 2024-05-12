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

func (r *ProductRepository) Store(p *ProductStorePayload) (*entities.Product, error) {
	log.Println(p)
	product := &entities.Product{}
	err := r.DB.QueryRow("INSERT INTO products (name, sku, category, image_url, notes, price, stock, location, is_available) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *", p.Name, p.SKU, p.Category, p.ImageUrl, p.Notes, p.Price, p.Stock, p.Location, p.IsAvailable).Scan(&product.ID, &product.Name, &product.SKU, &product.Category, &product.ImageUrl, &product.Notes, &product.Price, &product.Stock, &product.Location, &product.IsAvailable, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
	if err != nil {
		log.Printf("Error inserting product: %s", err)
		return nil, err
	}

	return product, nil
}
