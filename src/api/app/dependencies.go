package app

import (
	"github.com/IgancioRey/PedidosShaktyGo/src/api/config"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/controllers"
)

func InitDependencies() config.RoutesDependencies {
	routesDependencies := config.RoutesDependencies{
		GetProductsController:   controllers.GetProducts{},
		CreateProductController: controllers.CreateProduct{},
		GetCostumersController:  controllers.GetCustomers{},
	}

	routesDependencies.SetDependencies()

	return routesDependencies
}
