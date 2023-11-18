package database

import (
	"SportNotes/helper"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Host     = "localhost"
	Port     = "5432"
	User     = "postgres"
	Password = "postgres"
	DBName   = "sportnotes"
	SSLMode  = "disable"
)

// type Config struct {
// 	Host     string
// 	Port     string
// 	User     string
// 	Password string
// 	DBName   string
// 	SSLMode  string
// }

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		Host, Port, User, Password, DBName, SSLMode)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
