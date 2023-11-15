package routes

import (
	"SportNotes/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(workoutsController *controllers.WorkoutsController, accountsController *controllers.AccountsController) *gin.Engine {
	router := gin.Default()
	// Базовый роут
	baseRouter := router.Group("/api/v1")

	// Роуты для тренировок
	workoutsRouter := baseRouter.Group("/workouts")
	{
		// Получение всех записей о тренировках
		workoutsRouter.GET("", workoutsController.FindAll)

		// Создание записи тренировки
		workoutsRouter.POST("", workoutsController.Create)

		// Поиск записи тренировки по Id
		workoutsRouter.GET("/:workoutId", workoutsController.FindById)

		// Изменение записи тренировки
		// workoutsRouter.PATCH("/:workoutId", workoutsController.Update)

		// Удаление записи о тренировке
		workoutsRouter.DELETE("/:workoutId", workoutsController.Delete)
	}

	// Роуты для тренировок
	accountsRouter := baseRouter.Group("/accounts")
	{
		// Получение всех записей о тренировках
		accountsRouter.GET("", accountsController.FindAll)

		// Создание записи тренировки
		accountsRouter.POST("", accountsController.Create)

		// Поиск записи тренировки по Id
		accountsRouter.GET("/:accountId", accountsController.FindById)

		// Изменение записи тренировки
		// accountsRouter.PATCH("/:accountId", accountsController.Update)

		// Удаление записи о тренировке
		accountsRouter.DELETE("/:accountsId", accountsController.Delete)
	}
	return router
}
