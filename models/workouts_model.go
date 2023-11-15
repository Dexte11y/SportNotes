package models

import "time"

// Модель полей для таблицы workouts
type Workout struct {
	IdWorkout int       `gorm:"type:int"`
	IdAccount int       `gorm:"type:int"`
	Date      time.Time `gorm:"type:date"`
}
