package handlers

import (
	"cart/internal/core/domain"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (handler *HTTPHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	cart := domain.Cart{}
	newCart, err := handler.cartService.AddToCart(cart)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&newCart)
}

func (handler *HTTPHandler) DeleteAnItemFromCart(w http.ResponseWriter, r *http.Request) {
	cart := domain.Cart{}
	reference := chi.URLParam(r, "reference")
	deleteResponse, err := handler.cartService.DeleteAnItemFromCart(cart, reference)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(deleteResponse)
}

func (handler *HTTPHandler) DeleteAllCartItems(w http.ResponseWriter, r *http.Request) {
	cart := []domain.Cart{}
	deleteResponse, err := handler.cartService.DeleteAllCartItems(cart)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(deleteResponse)
}

func (handler *HTTPHandler) UpdateACartItem(w http.ResponseWriter, r *http.Request) {
	cart := domain.Cart{}
	reference := chi.URLParam(r, "reference")
	updateResponse, err := handler.cartService.UpdateACartItem(cart, reference)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(updateResponse)

}
// AddToCart(domain.Cart) (domain.Cart, error)
// DeleteAnItemFromCart(domain.Cart, string) (string, error)
// DeleteAllCartItems([]domain.Cart) (string, error)
// UpdateACartItem(domain.Cart, string) (string, error)
