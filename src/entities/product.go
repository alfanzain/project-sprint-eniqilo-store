package entities

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name" validate:"required,min=1,max=30"`
	SKU         string  `json:"sku" validate:"required,min=1,max=30"`
	Category    string  `json:"category" validate:"required"` // validate=enum`
	ImageUrl    string  `json:"imageUrl" validate:"required"`
	Notes       string  `json:"notes" validate:"required,min=1,max=200"`
	Price       int64   `json:"price" validate:"required,min=1"`
	Stock       int32   `json:"stock" validate:"required,min=0,max=100000"`
	Location    string  `json:"location" validate:"required,min=1,max=200"`
	IsAvailable bool    `json:"isAvailable" validate:"required"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	DeletedAt   *string `json:"deletedAt"`
}
