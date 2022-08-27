package config

import (
	"github.com/IgancioRey/PedidosShaktyGo/src/api/controllers"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/usecases"
)

type RoutesDependencies struct {
	GetProductsController   controllers.GetProducts
	GetProductController    controllers.GetProduct
	CreateProductController controllers.CreateProduct
	GetCustomersController  controllers.GetCustomers
	GetCustomerController   controllers.GetCustomer
}

func (rd *RoutesDependencies) SetDependencies() {
	getProductsUseCase := usecases.GetProductsImp{}
	getProductUseCase := usecases.GetProductImp{}
	createProductUseCase := usecases.CreateProductImp{}
	getCustomersUseCase := usecases.GetCustomersImp{}
	getCustomerUseCase := usecases.GetCustomerImp{}

	rd.GetProductsController.UseCase = &getProductsUseCase
	rd.GetProductController.UseCase = &getProductUseCase
	rd.CreateProductController.UseCase = &createProductUseCase
	rd.GetCustomersController.UseCase = &getCustomersUseCase
	rd.GetCustomerController.UseCase = &getCustomerUseCase

}
