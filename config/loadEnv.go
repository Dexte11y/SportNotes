package config

import (
	"SportNotes/helper"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	helper.ErrorPanic(err)
}
