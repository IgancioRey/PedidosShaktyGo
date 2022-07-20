package config

import (
	"github.com/IgancioRey/PedidosShaktyGo/src/api/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routers(gp controllers.GetProducts) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/products", gp.Execute)

	return router
}
