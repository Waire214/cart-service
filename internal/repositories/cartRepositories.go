package repositories

import (
	"cart/internal/core/domain"
	"cart/internal/core/ports/cartPort"
	"log"

	"github.com/jinzhu/gorm"
)

type cartInfra struct {
	db *gorm.DB
}

func NewCartRepositories(db *gorm.DB) cartPort.CartRepository {
	return &cartInfra{db}
}

func (r *cartInfra) AddToCart(cart domain.Cart) (domain.Cart, error) {
	newCart := domain.NewCart(cart.ProductId, cart.ProductName, cart.Quantity, cart.Price)

	if err := r.db.Create(newCart).Error; err != nil {
		log.Println(err)
	}
	return newCart, nil
}

func (r *cartInfra) DeleteAnItemFromCart(aCart domain.Cart, reference string) (string, error) {
	if err := r.db.Delete(&aCart).Error; err != nil {
		log.Println(err)
	}
	return reference, nil
}

func (r *cartInfra) DeleteAllCartItems(carts []domain.Cart) (string, error) {
	if err := r.db.Delete(&carts).Error; err != nil {
		log.Println(err)
	}
	return "empty", nil
}

func (r *cartInfra) UpdateACartItem(cart domain.Cart, reference string) (string, error) {
	// get the product price, insert later
	if err := r.db.Model(&cart).Where("reference = ?", reference).Select("Quantity", "Price").Updates(domain.Cart{Quantity: cart.Quantity - 1, Price: cart.Quantity * cart.Price}).Error; err != nil {
		log.Println(err)
	}

	return "updated", nil
}

// Select with Struct (select zero value fields)
// db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
// UPDATE users SET name='new_name', age=0 WHERE id=111;
// db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})

// AddToCart(domain.Cart) (domain.Cart, error)
// DeleteAnItemFromCart(domain.Cart, string) (string, error)
// DeleteAllCartItems([]domain.Cart) (string, error)
// UpdateACartItem(domain.Cart, string) (string, error)

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
