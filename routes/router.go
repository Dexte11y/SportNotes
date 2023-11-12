package routes

import (
	"SportNotes/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(workoutsController *controllers.WorkoutsController) *gin.Engine {
	router := gin.Default()

	baseRouter := router.Group("/api/v1")

	workoutsRouter := baseRouter.Group("/workouts")
	{
		workoutsRouter.GET("")
		workoutsRouter.POST("", workoutsController.Create)
	}
	return router
}
