package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	Reference   string    `json:"reference"`
	ProductId   string    `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func NewCart(productId, productName string, quantity, price int, createdAt, updatedAt, deletedAt time.Time) Cart {
	return Cart{
		Reference:   uuid.New().String(),
		ProductId:   productId,
		ProductName: productName,
		Quantity:    quantity,
		Price:       price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
