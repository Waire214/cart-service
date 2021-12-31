package main

import (
	"cart/helper"
	services "cart/internal/core/services/carts"
	"cart/internal/handlers"
	"cart/internal/repositories"
	"cart/routers"
	"fmt"
	"log"
	"net/http"
	"time"
)

func allConnection() {
	helper.InitializeLog()
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

		fmt.Println("App running on " + helper.Config.Address + ":" + helper.Config.Port)
		helper.LogEvent("Info", fmt.Sprintf("Started PlatformServiceApplication on "+helper.Config.Address+":"+helper.Config.Port+" in "+time.Since(time.Now()).String()))
		log.Fatal(http.ListenAndServe(":"+helper.Config.Port, helper.LogRequest(router)))
		// log.Fatal(http.ListenAndServe(":"+helper.Config.Port, router))
	}

}
