package main

import (
	"log"
	"net/http"

	"github.com/aharal/wager/api"
	"github.com/aharal/wager/api/v1/handlers"
	"github.com/aharal/wager/api/v1/middleware"
	"github.com/aharal/wager/api/v1/repository"
	"github.com/aharal/wager/api/v1/services"
	"github.com/aharal/wager/configs"
	"github.com/aharal/wager/database"
	"github.com/aharal/wager/database/connection"
	handler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//initialize the code
func init() {
	//Set Env variable for config file
	configs.Config.Read("production")
}

// We Manage / Set Environment in config.toml file
const (
	DEVELOPMENT = iota
	TEST
	STAGE
	PRODUCTION
)

func main() {
	database.DBMigrate() // DB Migration
	router := api.Route{}
	router.Router = mux.NewRouter()
	rm := middleware.RequestMiddleware{}
	router.Router.Use(rm.RequestIDGenerator)

	//Create Database Connection
	connectionService := connection.NewDatabaseConnection()
	connectionService.DBConnect()

	//Wagers
	wagersRepository := repository.NewWagerRepository(connectionService)
	wagersService := services.NewWagerService(wagersRepository)
	handlers.NewWagerHTTPHandler(wagersService, router)

	//BuyWagers
	buyWagersRepository := repository.NewBuyWagerRepository(connectionService)
	buyWagersService := services.NewBuyWagerService(buyWagersRepository)
	handlers.NewBuyWagerHTTPHandler(buyWagersService, router)

	log.Fatal(http.ListenAndServe(":"+configs.Config.Port, handler.CORS(handler.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handler.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handler.AllowedOrigins([]string{"*"}))(router.Router)))
}
