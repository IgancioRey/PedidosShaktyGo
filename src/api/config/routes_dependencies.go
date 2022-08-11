package config

import (
	"github.com/IgancioRey/PedidosShaktyGo/src/api/controllers"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/usecases"
)

type RoutesDependencies struct {
	GetProductsController   controllers.GetProducts
	GetProductController    controllers.GetProduct
	CreateProductController controllers.CreateProduct
	GetCostumersController  controllers.GetCustomers
}

func (rd *RoutesDependencies) SetDependencies() {
	getProductsUseCase := usecases.GetProductsImp{}
	getProductUseCase := usecases.GetProductImp{}
	createProductUseCase := usecases.CreateProductImp{}
	getCostumersUseCase := usecases.GetCustomersImp{}

	rd.GetProductsController.UseCase = &getProductsUseCase
	rd.GetProductController.UseCase = &getProductUseCase
	rd.CreateProductController.UseCase = &createProductUseCase
	rd.GetCostumersController.UseCase = &getCostumersUseCase

}
