package domain

import "github.com/google/uuid"

type Cart struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

func NewCart(name, state string) Cart {
	return Cart{
		Id: uuid.New().String(),
		Name: name,
		State: state,
	}
}
