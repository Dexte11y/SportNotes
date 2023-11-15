package repository

import (
	// "SportNotes/data/requests"
	"SportNotes/helper"
	"SportNotes/models"
	"errors"

	"gorm.io/gorm"
)

type WorkoutsRepositoryImpl struct {
	Db *gorm.DB
}

func NewWorkoutsRepositoryImpl(Db *gorm.DB) WorkoutsRepository {
	return &WorkoutsRepositoryImpl{Db: Db}
}

// Сохранение тренировок в БД
func (w *WorkoutsRepositoryImpl) Save(workouts models.Workout) {
	result := w.Db.Create(&workouts)
	helper.ErrorPanic(result.Error)

	// возвращает айди тренировки
	// fmt.Println(workouts.Id)
}

// Поиск всех тренировок из БД
func (w *WorkoutsRepositoryImpl) FindAll() []models.Workout {
	var Workouts []models.Workout
	result := w.Db.Find(&Workouts)
	helper.ErrorPanic(result.Error)
	return Workouts
}

// Поиск тренировки по Id в БД
func (w *WorkoutsRepositoryImpl) FindById(workoutsId int) (workouts models.Workout, err error) {
	var workout models.Workout
	result := w.Db.Find(&workout, workoutsId)
	if result != nil {
		return workout, nil
	} else {
		return workout, errors.New("workout is not found")
	}
}

// Обновление тренировки в БД
// func (w *WorkoutsRepositoryImpl) Update(workouts models.Workout) {
// 	var updateWorkouts = requests.UpdateWorkoutsRequest{
// 		Id:     workouts.Id,
// 		UserId: workouts.UserId,
// 		Date:   workouts.Date,
// 	}
// 	result := w.Db.Model(&workouts).Updates(updateWorkouts)
// 	helper.ErrorPanic(result.Error)
// }

// Удаление тренировки в БД
func (w *WorkoutsRepositoryImpl) Delete(workoutsId int) {
	var workouts models.Workout
	result := w.Db.Where("id = ?", workoutsId).Delete(&workouts)
	helper.ErrorPanic(result.Error)
}
