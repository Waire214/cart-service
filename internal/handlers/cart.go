package handlers

import (
	"cart/internal/core/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

func (handler *HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	cart := domain.Cart{}
	newCart, err := handler.cartService.Get(cart.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&newCart)
	// Get(id string) (domain.Cart, error)
	// Save(domain.Cart) (domain.Cart, error)

}

func (handler *HTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	cart := domain.Cart{}
	newCart, err := handler.cartService.Save(domain.NewCart(cart.Name, cart.State))
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(newCart)
}
