package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/config"
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

	routesDependencies := InitDependencies()

	Router := config.Routers(routesDependencies)

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = defaultPort
	}
	http.ListenAndServe(":"+serverPort, Router)
}
