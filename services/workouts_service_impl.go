package services

import (
	"SportNotes/helper"
	"SportNotes/models"
	"SportNotes/repository"
	"SportNotes/schemas"

	"github.com/go-playground/validator"
)

type WorkoutsServiceImpl struct {
	WorkoutsRepository repository.WorkoutsRepository
	Validate           *validator.Validate
}

func NewWorkoutsServiceImpl(WorkoutRepository repository.WorkoutsRepository, validate *validator.Validate) WorkoutsService {
	return &WorkoutsServiceImpl{
		WorkoutsRepository: WorkoutRepository,
		Validate:           validate,
	}
}

// Реализация сервиса тренировок
// Создание тренировки
func (w *WorkoutsServiceImpl) Create(workout schemas.CreateWorkoutSchema) {
	err := w.Validate.Struct(workout)
	helper.ErrorPanic(err)
	workoutModel := models.Workout{
		Id:        workout.Id,
		IdAccount: workout.IdAccount,
	}
	w.WorkoutsRepository.Save(workoutModel)
}

// Поиск всех тренировок
func (w *WorkoutsServiceImpl) FindAll() []schemas.ResponseWorkoutSchema {
	result := w.WorkoutsRepository.FindAll()

	var workouts []schemas.ResponseWorkoutSchema
	for _, value := range result {
		workout := schemas.ResponseWorkoutSchema{
			Id:        value.Id,
			IdAccount: value.IdAccount,
			CreatedAt: value.CreatedAt,
		}
		workouts = append(workouts, workout)
	}

	return workouts
}

// Обновление тренировки
// func (w *WorkoutsServiceImpl) Update(workouts requests.UpdateWorkoutsRequest) {
// workoutData, err := w.WorkoutsRepository.FindById(workouts.Id)
// helper.ErrorPanic(err)
// workoutData.Date = workouts.Date
// w.WorkoutsRepository.Update(workoutData)
// }

// Поиск тренировки по Id
func (w *WorkoutsServiceImpl) FindById(workoutsId int) schemas.ResponseWorkoutSchema {
	workoutData, err := w.WorkoutsRepository.FindById(workoutsId)
	helper.ErrorPanic(err)

	workoutResponse := schemas.ResponseWorkoutSchema{
		Id:        workoutData.Id,
		IdAccount: workoutData.IdAccount,
		CreatedAt: workoutData.CreatedAt,
	}
	return workoutResponse
}

// Удаление тренировки
func (w *WorkoutsServiceImpl) Delete(workoutsId int) {
	w.WorkoutsRepository.Delete(workoutsId)
}
