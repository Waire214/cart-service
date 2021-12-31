package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Cart struct {
	gorm.Model
	Reference   string `json:"reference"`
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity" validate:"min=8"`
	Price       int    `json:"price" validate:"min=2"`
}

func NewCart(productId, productName string, quantity, price int) *Cart {
	return &Cart{
		Reference:   uuid.New().String(),
		ProductId:   productId,
		ProductName: productName,
		Quantity:    quantity,
		Price:       price,
	}
}
