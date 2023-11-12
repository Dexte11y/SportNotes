package controllers

import (
	"SportNotes/services"

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

}
