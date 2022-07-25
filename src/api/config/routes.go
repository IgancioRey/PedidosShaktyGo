package config

import (
	"github.com/IgancioRey/PedidosShaktyGo/src/api/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routers(gp controllers.GetProducts, cp controllers.CreateProduct, gc controllers.GetCustomers) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/products", gp.Execute)
	router.Post("/products", cp.Execute)
	router.Get("/customers", gc.Execute)

	return router
}
