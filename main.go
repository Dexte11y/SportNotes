package main

import (
	"SportNotes/controllers"
	"SportNotes/database"
	"SportNotes/helper"
	"SportNotes/models"
	"SportNotes/repository"
	"SportNotes/routes"
	"SportNotes/services"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

func main() {
	log.Default().Println("Started Server!")
	// Database
	db := database.DatabaseConnection()
	validate := validator.New()

	db.Table("workouts").AutoMigrate(&models.Workouts{})

	// Repository
	WorkoutsRepository := repository.NewWorkoutsRepositoryImpl(db)

	// Service
	WorkoutsService := services.NewWorkoutsServiceImpl(WorkoutsRepository, validate)

	// Controller
	WorkoutsController := controllers.NewWorkoutsController(WorkoutsService)

	// Router
	routes := routes.NewRouter(WorkoutsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
