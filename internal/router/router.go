package router

import (
	"coffee_app_crud/internal/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/api/v1/coffees", controllers.GetAllCoffees)
	router.Get("/api/v1/coffees/{id}", controllers.GetCoffeeById)
	router.Delete("/api/v1/coffees/{id}", controllers.DeleteCoffee)
	router.Post("/api/v1/coffees/coffee", controllers.CreateCoffee)
	router.Put("/api/v1/coffees/{id}", controllers.UpdateCoffee)

	return router
}
