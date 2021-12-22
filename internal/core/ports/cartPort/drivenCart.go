package cartPort

import "cart/internal/core/domain"

type CartRepository interface {
	Get(id string) (domain.Cart, error)
	Save(domain.Cart) (domain.Cart, error)
}

// define the set of actions that the actor has to implement.
