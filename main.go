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

	db.Table("workout").AutoMigrate(&models.Workout{})
	db.Table("user").AutoMigrate(&models.Account{})
	// Repository
	WorkoutsRepository := repository.NewWorkoutsRepositoryImpl(db)
	AccountsRepository := repository.NewAccountsRepositoryImpl(db)
	// Service
	WorkoutsService := services.NewWorkoutsServiceImpl(WorkoutsRepository, validate)
	Accountservice := services.NewAccountsServiceImpl(AccountsRepository, validate)
	// Controller
	WorkoutsController := controllers.NewWorkoutsController(WorkoutsService)
	AccountsController := controllers.NewAccountsController(Accountservice)
	// Router
	routes := routes.NewRouter(WorkoutsController, AccountsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
