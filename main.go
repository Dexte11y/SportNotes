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
	// db, err := database.DatabaseConnection(database.Config{
	// 	Host:     viper.GetString("db.host"),
	// 	Port:     viper.GetString("db.port"),
	// 	User:     viper.GetString("db.user"),
	// 	DBName:   viper.GetString("db.dbname"),
	// 	Password: viper.GetString("db.password"),
	// 	SSLMode:  viper.GetString("db.sslmode"),
	// })
	// if err != nil {
	// 	log.Fatal("failed to initialize db:", err.Error())
	// }

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

// func initConfig() error {
// 	viper.AddConfigPath("configs")
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")

// 	err := viper.ReadInConfig()
// 	helper.ErrorPanic(err)

// 	return nil
// }
