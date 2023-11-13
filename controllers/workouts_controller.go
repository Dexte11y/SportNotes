package controllers

import (
	"SportNotes/data/requests"
	"SportNotes/data/responses"
	"SportNotes/helper"
	"SportNotes/services"
	"log"
	"net/http"
	"strconv"

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

// Контроллер на создание тренировки
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

// Контроллер на поиск тренировок
func (controller *WorkoutsController) FindAll(c *gin.Context) {
	log.Default().Println("FindAll workouts")
	workoutResponse := controller.workoutsService.FindAll()
	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   workoutResponse,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// Контроллер на обновление тренировки
// func (controller *WorkoutsController) Update(c *gin.Context) {
// 	log.Default().Println("Update workouts")
// 	updateWorkoutsRequest := requests.UpdateWorkoutsRequest{}
// 	err := c.BindJSON(&updateWorkoutsRequest)
// 	helper.ErrorPanic(err)

// 	workoutId := c.Param("workoutId")
// 	id, err := strconv.Atoi(workoutId)
// 	helper.ErrorPanic(err)
// 	updateWorkoutsRequest.Id = id

// 	controller.workoutsService.Update(updateWorkoutsRequest)

// 	webResponse := responses.Response{
// 		Code:   http.StatusOK,
// 		Status: "Ok",
// 		Data:   nil,
// 	}
// 	c.Header("Content-Type", "application/json")
// 	c.JSON(http.StatusOK, webResponse)
// }

// Контроллер на поиск тренировки по Id
func (controller *WorkoutsController) FindById(c *gin.Context) {
	log.Default().Println("FindById workouts")
	workoutId := c.Param("workoutId")
	id, err := strconv.Atoi(workoutId)
	helper.ErrorPanic(err)

	workoutResponse := controller.workoutsService.FindById(id)

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   workoutResponse,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// Контроллер на удаление тренировки
func (controller *WorkoutsController) Delete(c *gin.Context) {
	log.Default().Println("Delete workouts")
	workoutId := c.Param("workoutId")
	id, err := strconv.Atoi(workoutId)
	helper.ErrorPanic(err)
	controller.workoutsService.Delete(id)

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
