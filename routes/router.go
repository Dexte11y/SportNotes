package routes

import (
	"SportNotes/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(workoutsController *controllers.WorkoutsController) *gin.Engine {
	router := gin.Default()
	// Базовый роут
	baseRouter := router.Group("/api/v1")

	// Роуты для тренировок
	workoutsRouter := baseRouter.Group("/workouts")
	{
		// Получение всех тренировок
		workoutsRouter.GET("", workoutsController.FindAll)
		// Создание записи тренировки
		workoutsRouter.POST("", workoutsController.Create)
	}
	return router
}
