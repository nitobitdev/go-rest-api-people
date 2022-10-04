package handlers

import (
	"go-rest-api-people/controllers"

	docs "go-rest-api-people/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	people *controllers.PersonController
}

func NewRouter(people *controllers.PersonController) *Router {
	return &Router{people: people}
}

func (r *Router) Start(port string) {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/people", r.people.GetPeople)
	router.GET("/people/:id", r.people.GetPeopleById)
	router.POST("/people", r.people.CreatePeople)
	router.PUT("/people/:id", r.people.UpdatePeople)
	router.DELETE("/people/:id", r.people.DeletePeople)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(port)
}
