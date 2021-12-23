package handlers

import "cart/internal/core/ports/cartPort"

type HTTPHandler struct {
	cartService cartPort.CartService
}

func NewHTTPHandler(cartService cartPort.CartService) *HTTPHandler {
	return &HTTPHandler{
		cartService: cartService,
	}
}

func (handler *HTTPHandler) Get() {
	// Get(id string) (domain.Cart, error)
	// Save(domain.Cart) (domain.Cart, error)

}