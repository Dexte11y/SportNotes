package repository

import (
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
func (w *WorkoutsRepositoryImpl) Save(workout models.Workout) {
	result := w.Db.Table("workout").Create(&workout)
	helper.ErrorPanic(result.Error)
}

// Поиск всех тренировок из БД
func (w *WorkoutsRepositoryImpl) FindAll() []models.Workout {
	var Workouts []models.Workout
	result := w.Db.Table("workout").Find(&Workouts)
	helper.ErrorPanic(result.Error)
	return Workouts
}

// Поиск тренировки по Id в БД
func (w *WorkoutsRepositoryImpl) FindById(workoutsId int) (workouts models.Workout, err error) {
	var workout models.Workout
	result := w.Db.Table("workout").Find(&workout, workoutsId)
	if result != nil {
		return workout, nil
	} else {
		return workout, errors.New("тренировка не найдена")
	}
}

// Обновление тренировки в БД
// func (w *WorkoutsRepositoryImpl) Update(workouts models.Workout) {
// 	var updateWorkouts = requests.UpdateWorkoutsRequest{
// 		Id:     workouts.Id,
// 		UserId: workouts.UserId,
// 		Date:   workouts.Date,
// 	}
// result := w.Db.Table("workout").Model(&workout).Updates(updateWorkouts)
// 	helper.ErrorPanic(result.Error)
// }

// Удаление тренировки в БД
func (w *WorkoutsRepositoryImpl) Delete(workoutId int) {
	var workout models.Workout
	result := w.Db.Table("workout").Where("id = ?", workoutId).Delete(&workout)
	helper.ErrorPanic(result.Error)
}
