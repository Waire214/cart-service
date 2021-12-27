package handlers

import (
	"cart/internal/core/ports/cartPort"
)

type HTTPHandler struct {
	cartService cartPort.CartService
}

func NewHTTPHandler(cartService cartPort.CartService) *HTTPHandler {
	return &HTTPHandler{
		cartService: cartService,
	}
}
