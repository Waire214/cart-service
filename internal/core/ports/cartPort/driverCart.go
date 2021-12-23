package cartPort

import "cart/internal/core/domain"

type CartService interface {
	Get(id string) (domain.Cart, error)
	Save(domain.Cart) (domain.Cart, error)
}
