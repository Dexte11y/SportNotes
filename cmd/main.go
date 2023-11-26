package main

import (
	"SportNotes/config"
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

	config.LoadEnv()

	// Database
	db := database.DatabaseConnection()

	validate := validator.New()

	db.Table("account").AutoMigrate(&models.Account{})
	db.Table("workout").AutoMigrate(&models.Workout{})
	db.Table("training").AutoMigrate(models.Training{})

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
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
