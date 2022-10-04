package controllers

import (
	"go-rest-api-people/models"
	"go-rest-api-people/views"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonController struct {
	db *gorm.DB
}

func NewPersonController(db *gorm.DB) *PersonController {
	return &PersonController{
		db: db,
	}
}

// GetPeoples godoc
// @Summary Get all people
// @Decription Get list people
// @Tags People
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /peopels [get]
func (p *PersonController) GetPeople(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})

}

// GetPeopleById          godoc
// @Summary      Get People By Id
// @Description  Get a People.
// @Tags         People
// @Produce      json
// @Param 			id	path int true "Get People By Id"
// @Success      200   {object}  models.Person
// @Router       /peoples/{id} [get]
func (p *PersonController) GetPeopleById(ctx *gin.Context) {
	people := models.Person{}

	id, _ := strconv.Atoi(ctx.Param("id"))
	err := p.db.Where("id = ?", id).First(&people).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, &views.Response{
			Status:  http.StatusOK,
			Message: "GET_PEOPLE_SUCCESS",
			Payload: people,
		})
	}
}

// PostPeople          godoc
// @Summary      Store a new people
// @Description  Insert a person JSON and store in DB. Return saved JSON.
// @Tags         People
// @Produce      json
// @Param        people body models.RequestPerson true "Insert People"
// @Success      201   {object}  models.Person
// @Router       /peoples [post]
func (p *PersonController) CreatePeople(ctx *gin.Context) {
	people := models.Person{}

	people.Name = ctx.PostForm("name")
	people.Address = ctx.PostForm("address")

	p.db.Create(&people)
	ctx.JSON(http.StatusCreated, &views.Response{
		Status:  http.StatusCreated,
		Message: "CREATE_PEOPLE_SUCCESS",
		Payload: people,
	})

}

// UpdatePeople          godoc
// @Summary      Update a people
// @Description  Update a person JSON and store in DB. Return saved JSON.
// @Tags         People
// @Produce      json
// @Param 			 id path int true "Update People By Id"
// @Param        people body models.RequestPerson true "Update People"
// @Success      201   string  "Success Update People"
// @Router       /peoples{id} [put]
func (p *PersonController) UpdatePeople(ctx *gin.Context) {
	id := ctx.Param("id")

	people := models.Person{}
	newPeople := models.Person{}

	err := p.db.Find(&people, id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}

	newPeople.Name = ctx.PostForm("name")
	newPeople.Address = ctx.PostForm("address")

	err = p.db.Model(&people).Updates(newPeople).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, &views.Response{
			Status:  http.StatusOK,
			Message: "SUCCESS_UPDATED_DATA_PEOPLE",
		})
		return
	}
}

// DeletePopleById          godoc
// @Summary      Delete People By Id
// @Description  Delete a People.
// @Tags         People
// @Produce      json
// @Param 			id	path int true "Delete People By Id"
// @Success      200   string  "Success Delete People"
// @Router       /peoples/{id} [delete]
func (p *PersonController) DeletePeople(ctx *gin.Context) {
	people := models.Person{}

	id := ctx.Param("id")
	err := p.db.First(&people, id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
	}

	err = p.db.Delete(&people).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, &views.Response{
			Status:  http.StatusOK,
			Message: "DELETE_PEOPLE_SUCCESS",
		})
		return
	}
}
