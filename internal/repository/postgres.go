package repository

import (
	"fmt"
	"sportnotes/configs"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	workoutsTable   = "workouts"
	activityTable   = "activity"
	nutritionsTable = "nutritions"
	foodsTable      = "foods"
)

func NewPostgresDB(cfg configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.DBName, cfg.DB.Password, cfg.DB.SSLMode))
	if err != nil {
		return nil, err
	}

	return db, nil
}
