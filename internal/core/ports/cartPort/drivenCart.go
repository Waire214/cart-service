package cartPort

import "cart/internal/core/domain"

type CartRepository interface {
	AddToCart(domain.Cart) (domain.Cart, error)
	DeleteAnItemFromCart(reference string) (string, error)
	DeleteAllCartItems(domain.Cart) (string, error)
	UpdateACartItem(reference string) (string, error)
}

// define the set of actions that the actor has to implement.
