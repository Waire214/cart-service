package main

import (
	"cart/helper"
	services "cart/internal/core/services/carts"
	"cart/internal/handlers"
	"cart/internal/repositories"
	"cart/routers"
	"net/http"
)

func allConnection() {

	// database connection
	db, err := repositories.DataBaseConnection(helper.Config.Mode, helper.Config.DbHost, helper.Config.User, helper.Config.DbName, helper.Config.Password)
	if err != nil {
		helper.PrintErrorMessage("404", err.Error())
	}
	//repositories connection
	cartRepository := repositories.NewCartRepositories(db)

	//services connection
	cartService := services.NewCartService(cartRepository)

	//handlers connection
	handler := handlers.NewHTTPHandler(cartService)

	//routers connection
	if helper.Config.Mode == "dev" {
		router := routers.Router(helper.Config.Port, helper.Config.Address, handler)
		http.ListenAndServe(":"+helper.Config.Port, router)
	}

}
