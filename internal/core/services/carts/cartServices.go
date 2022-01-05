package services

import (
	"cart/helper"
	"cart/internal/core/domain"
	"cart/internal/core/ports/cartPort"
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

	if err := helper.ValidateStruct(cart); err != nil {
		return domain.Cart{}, err
	}
	return serve.cartRepository.AddToCart(cart)
}

func (serve *cartService) DeleteAnItemFromCart(cart domain.Cart, reference string) (string, error) {
	return serve.cartRepository.DeleteAnItemFromCart(cart, reference)
}

func (serve *cartService) DeleteAllCartItems(carts []domain.Cart) (string, error) {
	return serve.cartRepository.DeleteAllCartItems(carts)
}

func (serve *cartService) ReduceQuantity(cart domain.Cart, reference string) (string, error) {
	return serve.cartRepository.ReduceQuantity(cart, reference)

}
func (serve *cartService) IncreaseQuantity(cart domain.Cart, reference string) (string, error) {
	return serve.cartRepository.IncreaseQuantity(cart, reference)
}
// AddToCart(domain.Cart) (domain.Cart, error)
// DeleteAnItemFromCart(domain.Cart, string) (string, error)
// DeleteAllCartItems([]domain.Cart) (string, error)
// UpdateACartItem(domain.Cart, string) (string, error)
