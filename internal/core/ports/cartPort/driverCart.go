package cartPort

import "cart/internal/core/domain"

type CartService interface {
	AddToCart(domain.Cart) (domain.Cart, error)
	DeleteAnItemFromCart(domain.Cart, string) (string, error)
	DeleteAllCartItems([]domain.Cart) (string, error)
	ReduceQuantity(domain.Cart, string) (string, error)
	IncreaseQuantity(domain.Cart, string) (string, error)
}
