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

func (serve *cartService) AddToCart(cart domain.Cart) (domain.Cart, error) {
	cart, err := serve.cartRepository.AddToCart(cart)
	if err != nil {
		return domain.Cart{}, errors.New("err")
	}
	return cart, nil
}

func (serve *cartService) DeleteAnItemFromCart(cart domain.Cart,reference string) (string, error) {
	deleteResponse, err := serve.cartRepository.DeleteAnItemFromCart(cart, reference)
	if err != nil {
		return err.Error(), errors.New("err")
	}
	return deleteResponse, nil
}

func (serve *cartService) DeleteAllCartItems(carts []domain.Cart) (string, error) {
	deleteResponse, err := serve.cartRepository.DeleteAllCartItems(carts)
	if err != nil {
		return err.Error(), errors.New("err")
	}
	return deleteResponse, nil
}

func (serve *cartService) UpdateACartItem(cart domain.Cart, reference string) (string, error) {
	updateResponse, err := serve.cartRepository.UpdateACartItem(cart, reference)
	if err != nil {
		return err.Error(), errors.New("err")
	}
	return updateResponse, nil

}
// AddToCart(domain.Cart) (domain.Cart, error)
// DeleteAnItemFromCart(domain.Cart, string) (string, error)
// DeleteAllCartItems([]domain.Cart) (string, error)
// UpdateACartItem(domain.Cart, string) (string, error)
