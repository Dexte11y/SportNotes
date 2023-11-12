package controllers

import (
	"SportNotes/data/requests"
	"SportNotes/data/responses"
	"SportNotes/helper"
	"SportNotes/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorkoutsController struct {
	workoutsService services.WorkoutsService
}

func NewWorkoutsController(service services.WorkoutsService) *WorkoutsController {
	return &WorkoutsController{
		workoutsService: service,
	}
}

func (controller *WorkoutsController) Create(c *gin.Context) {
	log.Default().Println("Create workouts")
	createWorkoutsRequest := requests.CreateWorkoutsRequest{}
	err := c.BindJSON(&createWorkoutsRequest)
	helper.ErrorPanic(err)

	controller.workoutsService.Create(createWorkoutsRequest)
	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
