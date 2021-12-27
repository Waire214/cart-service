package repositories

import (
	"cart/internal/core/domain"
	"cart/internal/core/ports/cartPort"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type cartInfra struct {
	db *gorm.DB
}

func NewCartRepositories(db *gorm.DB) cartPort.CartRepository {
	return &cartInfra{db}
}

func (r *cartInfra) AddToCart(cart domain.Cart) (domain.Cart, error) {
	newCart := domain.NewCart(cart.ProductId, cart.ProductName, cart.Quantity, cart.Price, time.Now(), time.Time{}, time.Time{})

	if err := r.db.Create(newCart).Error; err != nil {
		log.Println(err)
	}
	return newCart, nil
}

func (r *cartInfra) DeleteAnItemFromCart(reference string) (domain.Cart, error) {
	r.db.Create(domain.NewCart("Ella", "Lagos"))
	return domain.NewCart("Ella", "Lagos"), nil
}

// AddToCart(domain.Cart) (domain.Cart, error)
// DeleteAnItemFromCart(reference string) (string, error)
// DeleteAllCartItems(domain.Cart) (string, error)
// UpdateACartItem(reference string) (string, error)
// productId string, productName string, quantity int, price int, createdAt time.Time, updatedAt time.Time, deletedAt time.Time

// You can find soft deleted records with Unscoped
// db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;

// You can delete matched records permanently with Unscoped
// db.Unscoped().Delete(&order)
// DELETE FROM orders WHERE id=10;

// db.Delete(&user)
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

// Batch Delete
// db.Where("age = ?", 20).Delete(&User{})
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// Soft deleted records will be ignored when querying
// db.Where("age = 20").Find(&user)
// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
