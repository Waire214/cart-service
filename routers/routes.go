package routers

import (
	"cart/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Router(appPort, hostAddress string, handler *handlers.HTTPHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// router.Use(middleware.Logger)
	// router.Use(helper.LogRequest)
	router.Mount("/cart", cartEndpoint(handler))
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(hostAddress+":"+appPort+"/swagger/doc.json"),
	))

	return router
}

func cartEndpoint(handler *handlers.HTTPHandler) http.Handler {
	router := chi.NewRouter()
	// router.Use(helper.LogRequest)
	router.Post("/add", handler.AddToCart)
	router.Delete("/delete/1/{reference}", handler.DeleteAnItemFromCart)
	router.Delete("/delete/all", handler.DeleteAllCartItems)
	router.Put("/update/reduce/1", handler.ReduceQuantity)
	router.Put("/update/increase/1", handler.IncreaseQuantity)
	return router
}
