package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/config"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/controllers"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/usecases"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/repository/database"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func Start() {
	fmt.Println("Starting Pedidos Shakty application")
	database.SetupDatabase()
	initRouter()
}

func initRouter() {

	getProductsUseCase := usecases.GetProductsImp{}
	getProductsController := controllers.GetProducts{
		UseCase: &getProductsUseCase,
	}

	createProductUseCase := usecases.CreateProductImp{}
	createProductController := controllers.CreateProduct{
		UseCase: &createProductUseCase,
	}

	getCostumersUseCase := usecases.GetCustomersImp{}
	getCostumersController := controllers.GetCustomers{
		UseCase: &getCostumersUseCase,
	}

	Router := config.Routers(getProductsController, createProductController, getCostumersController)

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = defaultPort
	}
	http.ListenAndServe(":"+serverPort, Router)
}
