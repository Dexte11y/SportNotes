package repository

import (
	"fmt"
	"sportnotes/configs"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	workoutsTable   = "workouts"
	trainingsTable  = "trainings"
	nutritionsTable = "nutritions"
	foodsTable      = "foods"
)

func NewPostgresDB(cfg configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Db.Host, cfg.Db.Port, cfg.Db.Username, cfg.Db.DBName, cfg.Db.Password, cfg.Db.SSLMode))
	if err != nil {
		return nil, err
	}

	return db, nil
}
