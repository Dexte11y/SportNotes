package handler

import (
	"sportnotes/pkg/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "sportnotes/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.Default())

	// auth := router.Group("/auth")
	// {
	// 	auth.POST("/sign-up", h.signUp)
	// 	auth.POST("/sign-in", h.signIn)
	// }

	api := router.Group("/api")
	{
		food := api.Group("/foods")
		{
			food.POST("/", h.createFood)
			food.GET("/", h.getAllFoods)
			food.GET("/:id", h.getFoodByID)
			// food.PUT("/:id", h.updateFood)
			food.DELETE("/:id", h.deleteFood)
		}

		nutritions := api.Group("/nutritions")
		{
			nutritions.POST("/:idUser", h.createNutrition)
			nutritions.GET("/:id", h.getNutritionsByParam)
			// nutritions.PUT("/:id", h.updateNutrition)
			nutritions.DELETE("/:id", h.deleteNutrition)
		}

		trainings := api.Group("/activity")
		{
			trainings.POST("/", h.createActivity)
			trainings.GET("/", h.getAllActivity)
			trainings.GET("/:id", h.getActivityByID)
			trainings.PUT("/:id", h.updateActivity)
			trainings.DELETE("/:id", h.deleteActivity)
		}

		workouts := api.Group("/workouts")
		{
			workouts.POST("/:idUser", h.createWorkout)
			workouts.GET("/:id", h.getWorkoutsByParam)
			// workouts.PUT("/:id", h.updateWorkout)
			workouts.DELETE("/delete/:id", h.deleteWorkout)
		}

		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserByID)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}
	}

	return router
}
