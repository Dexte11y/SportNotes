package handler

import (
	"sportnotes/pkg/service"

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

	// auth := router.Group("/auth")
	// {
	// 	auth.POST("/sign-up", h.signUp)
	// 	auth.POST("/sign-in", h.signIn)
	// }

	// api := router.Group("/api", h.doctorIdentity)
	// , h.userIdentity
	api := router.Group("/api")
	{
		// doctors := api.Group("/doctors")
		// {
		// 	doctors.GET("/", h.getAllDoctors)
		// 	doctors.GET("/:id", h.getDoctorById)
		// }

		food := api.Group("/foods")
		{
			food.POST("/", h.createFood)
			food.GET("/", h.getAllFoods)
			food.GET("/:id", h.getFoodById)
			// food.PUT("/:id", h.updateFood)
			food.DELETE("/:id", h.deleteFood)
		}

		nutritions := api.Group("/nutritions")
		{
			nutritions.POST("/", h.createNutrition)
			nutritions.GET("/:id", h.getNutritionsByParam)
			// nutritions.PUT("/:id", h.updateNutrition)
			nutritions.DELETE("/:id", h.deleteNutrition)
		}

		trainings := api.Group("/trainings")
		{
			trainings.POST("/", h.createTraining)
			trainings.GET("/", h.getAllTrainings)
			trainings.GET("/:id", h.getTrainingById)
			trainings.PUT("/:id", h.updateTraining)
			trainings.DELETE("/:id", h.deleteTraining)
		}

		workouts := api.Group("/workouts")
		{
			workouts.POST("/", h.createWorkout)
			workouts.GET("/:id", h.getWorkoutsByParam)
			// workouts.PUT("/:id", h.updateWorkout)
			workouts.DELETE("/:id", h.deleteWorkout)
		}

		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)

		}
	}

	return router
}
