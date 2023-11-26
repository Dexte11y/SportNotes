package database

import (
	"SportNotes/helper"
	"SportNotes/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {

	dbHost := utils.DefaultGetEnv("HOST", "")
	dbPort := utils.DefaultGetEnv("PORT", "")
	dbUser := utils.DefaultGetEnv("USER", "")
	dbPassword := utils.DefaultGetEnv("PASSWORD", "")
	dbName := utils.DefaultGetEnv("DBNAME", "")
	dbSSLMode := utils.DefaultGetEnv("SSLMODE", "")

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
