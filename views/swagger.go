package views

import "go-rest-api-people/models"

type GetAllPeopleSwagger struct {
	Response
	Payload []models.Person
}
