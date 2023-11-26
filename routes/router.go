package routes

import (
	"SportNotes/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(workoutsController *controllers.WorkoutsController, accountsController *controllers.AccountsController) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	// Базовый роут
	baseRouter := router.Group("/api/v1")

	// Роут для главной страницы
	// homePage := baseRouter.Group("/home")
	// {
	// 	homePage.GET("", homePageController.GetHomePage)

	// }

	temp := baseRouter.Group("/temp")
	{
		temp.GET("/home", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", nil)
		})

		temp.GET("/auth", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "auth.html", nil)
		})

		temp.GET("/workouts", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "workout.html", nil)
		})
	}

	// authRouter := baseRouter.Group("/auth")
	// {
	// authRouter.POST("/login", accountsController.Login)
	// authRouter.POST("/register", accountsController.Register)
	// }

	// traningRouter := baseRouter.Group("/tranings")
	// {
	// traningRouter.GET("", workoutsController.FindAll)
	// traningRouter.POST("", workoutsController.Create)
	// }

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

	// Роуты для аккаунтов
	accountsRouter := baseRouter.Group("/accounts")
	{
		// Получение всех аккаунтов
		accountsRouter.GET("", accountsController.FindAll)

		// Регистрация аккаунта
		accountsRouter.POST("/register", accountsController.Create)

		// Авторизация аккаунта
		accountsRouter.POST("/login", accountsController.Login)
		
		// Поиск аккаунта по Id
		accountsRouter.GET("/:accountId", accountsController.FindById)

		// Изменение аккаунта
		// accountsRouter.PATCH("/:accountId", accountsController.Update)

		// Удаление аккаунта
		accountsRouter.DELETE("/:accountsId", accountsController.Delete)
	}

	return router
}
