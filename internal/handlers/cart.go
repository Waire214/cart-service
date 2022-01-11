package handlers

import (
	"cart/internal/core/domain"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//AddToCart godoc
//@Summary handles adding items to cart
//@Description Accept JSON data of cart data and returns valid response
//@Accept json
//@Tags Cart
//@Produce json
//@Param Cart body domain.Cart true "Add items to cart"
//@Param authorization header string true "Authentication Token"
//@Success 200 {object} domain.Cart "Item added to cart"
//@Router /cart/add [post]
func (handler *HTTPHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	cart := domain.Cart{}
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	newCart, err := handler.cartService.AddToCart(cart)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(newCart)
}

//DeleteAnItemFromCart godoc
//@Summary handles deleting an item from cart
//@Description Take the item-reference as a PARAMETER and returns valid response
//@Accept json
//@Tags Cart
//@Produce json
//@Success 200 {object} domain.Cart "deleted"
//@Router /cart/delete/1/{reference} [delete]
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

//DeleteAllCartItems godoc
//@Summary handles deleting all item from cart
//@Description deletes all items from cart and returns valid response
//@Accept json
//@Tags Cart
//@Produce json
//@Success 200 {object} domain.Cart.Reference "all cart items deleted"
//@Router /cart/delete/all [delete]
func (handler *HTTPHandler) DeleteAllCartItems(w http.ResponseWriter, r *http.Request) {
	cart := []domain.Cart{}
	deleteResponse, err := handler.cartService.DeleteAllCartItems(cart)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(deleteResponse)
}

//ReduceQuantity godoc
//@Summary handles updating items in cart
//@Description Accept JSON data of cart data, takes the reference as a parameter to query the database and returns valid response
//@Accept json
//@Tags Cart
//@Produce json
//@Param Cart body domain.Cart true "Update cart items"
//@Param authorization header string true "Authentication Token"
//@Success 200 {object} domain.Cart "Item Updated"
//@Router /cart/update/reduce/1 [put]
func (handler *HTTPHandler) ReduceQuantity(w http.ResponseWriter, r *http.Request) {
	cart := domain.Cart{}
	reference := chi.URLParam(r, "reference")
	updateResponse, err := handler.cartService.ReduceQuantity(cart, reference)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(updateResponse)

}

//IncreaseQuantity godoc
//@Summary handles updating items in cart
//@Description Accept JSON data of cart data, takes the reference as a parameter to query the database and returns valid response
//@Accept json
//@Tags Cart
//@Produce json
//@Param Cart body domain.Cart true "Update cart items"
//@Param authorization header string true "Authentication Token"
//@Success 200 {object} domain.Cart "Item Updated"
//@Router /cart/update/increase/1 [put]
func (handler *HTTPHandler) IncreaseQuantity(w http.ResponseWriter, r *http.Request) {
	cart := domain.Cart{}
	reference := chi.URLParam(r, "reference")
	updateResponse, err := handler.cartService.IncreaseQuantity(cart, reference)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(updateResponse)

}
