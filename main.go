package main

import (
	"go-rest-api-people/controllers"
	"go-rest-api-people/handlers"

	"go-rest-api-people/config"
)

func main() {

	db := config.ConnectGorm()
	peopleController := controllers.NewPersonController(db)
	router := handlers.NewRouter(peopleController)
	router.Start(":4000")

}
