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
