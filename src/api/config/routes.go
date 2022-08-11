package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routers(rd RoutesDependencies) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/products", rd.GetProductsController.Execute)
	router.Get("/products/{id}", rd.GetProductController.Execute)
	router.Post("/products", rd.CreateProductController.Execute)
	router.Get("/customers", rd.GetCostumersController.Execute)

	return router
}
