package services

import (
	"cart/internal/core/domain"
	"cart/internal/core/ports/cartPort"
	"errors"
)

type cartService struct {
	cartRepository cartPort.CartRepository
}

func NewCartService(cartRepository cartPort.CartRepository) *cartService {
	return &cartService{
		cartRepository: cartRepository,
	}
}

func (serve *cartService) Get(id string) (domain.Cart, error) {
	cart, err := serve.cartRepository.Get(id)
	if err != nil {
		return domain.Cart{}, errors.New("err")
	}
	return cart, nil
}

func (serve *cartService) Save(cart domain.Cart) (domain.Cart,error) {
	newCart, err := serve.cartRepository.Save(domain.NewCart(cart.Name, cart.State))
	if err != nil {
		return domain.Cart{}, errors.New("err")
	}
	return newCart, nil
}