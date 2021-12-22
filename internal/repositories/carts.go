package repositories

import (
	"cart/internal/core/domain"
	"cart/internal/core/ports/cartPort"

	"github.com/jinzhu/gorm"
)

type cartInfra struct {
	db *gorm.DB
}

func NewcartInfra(db *gorm.DB) cartPort.CartRepository {
	return &cartInfra{db}
}


func (r *cartInfra) Get(id string) (domain.Cart, error) {
	return domain.NewCart("Ella", "Lagos"), nil

}

func (r *cartInfra) Save(domain.Cart) (domain.Cart,error) {
	r.db.Create(domain.NewCart("Ella", "Lagos"))
	return domain.NewCart("Ella", "Lagos"), nil
}