package config

import (
	"github.com/IgancioRey/PedidosShaktyGo/src/api/controllers"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/usecases"
)

type RoutesDependencies struct {
	GetProductsController   controllers.GetProducts
	CreateProductController controllers.CreateProduct
	GetCostumersController  controllers.GetCustomers
}

func (rd *RoutesDependencies) SetDependencies() {
	getProductsUseCase := usecases.GetProductsImp{}
	createProductUseCase := usecases.CreateProductImp{}
	getCostumersUseCase := usecases.GetCustomersImp{}

	rd.GetProductsController.UseCase = &getProductsUseCase
	rd.CreateProductController.UseCase = &createProductUseCase
	rd.GetCostumersController.UseCase = &getCostumersUseCase

}
