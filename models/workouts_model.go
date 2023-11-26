package models

import "time"

// Модель полей для таблицы workout
type Workout struct {
	Id        int       `gorm:"type:int;primary_key"`
	IdAccount int       `gorm:"type:int"`
	CreatedAt time.Time `gorm:"type:date"`
}
