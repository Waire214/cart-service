package cartPort

import "cart/internal/core/domain"

type CartRepository interface {
	AddToCart(domain.Cart) (domain.Cart, error)
	DeleteAnItemFromCart(domain.Cart, string) (string, error)
	DeleteAllCartItems([]domain.Cart) (string, error)
	ReduceQuantity(domain.Cart, string) (string, error)
	IncreaseQuantity(domain.Cart, string) (string, error)
}

// define the set of actions that the actor has to implement.
