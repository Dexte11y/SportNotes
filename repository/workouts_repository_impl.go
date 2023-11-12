package repository

import (
	"SportNotes/helper"
	"SportNotes/models"
	"fmt"

	"gorm.io/gorm"
)

type WorkoutsRepositoryImpl struct {
	Db *gorm.DB
}

func NewWorkoutsRepositoryImpl(Db *gorm.DB) WorkoutsRepository {
	return &WorkoutsRepositoryImpl{Db: Db}
}

// Сохранение тренировок в БД
func (w *WorkoutsRepositoryImpl) Save(workouts models.Workouts) {
	result := w.Db.Create(&workouts)
	helper.ErrorPanic(result.Error)

	// возвращает айди тренировки
	fmt.Println(workouts.Id)
}

// Поиск всех тренировок из БД
func (w *WorkoutsRepositoryImpl) FindAll() []models.Workouts {
	var Workouts []models.Workouts
	result := w.Db.Find(&Workouts)
	helper.ErrorPanic(result.Error)
	return Workouts
}
