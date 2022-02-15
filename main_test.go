package main

import (
	"testing"

	"github.com/aharal/wager/api"
	"github.com/aharal/wager/api/v1/handlers"
	"github.com/aharal/wager/api/v1/repository"
	"github.com/aharal/wager/api/v1/services"
	"github.com/aharal/wager/configs"
	"github.com/aharal/wager/database"
	"github.com/aharal/wager/database/connection"
	"github.com/gorilla/mux"
)

//initialize the code
func init() {
	//Set Env variable for config file
	configs.Config.Read("testing")
}

func Testmain(m *testing.M) {
	database.DBMigrate() // DB Migration
	//Create Database Connection
	fakeConnectionService := connection.NewDatabaseConnection()
	fakeConnectionService.DBConnect()
	router := api.Route{}
	router.Router = mux.NewRouter()

	//Wager
	albumsRepository := repository.NewWagerRepository(fakeConnectionService)
	albumsService := services.NewWagerService(albumsRepository)
	handlers.NewWagerHTTPHandler(albumsService, router)

	m.Run()
}
