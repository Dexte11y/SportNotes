package services

import (
	"SportNotes/data/requests"
	"SportNotes/data/responses"
	"SportNotes/helper"
	"SportNotes/models"
	"SportNotes/repository"

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
func (w *WorkoutsServiceImpl) Create(workout requests.CreateWorkoutsRequest) {
	err := w.Validate.Struct(workout)
	helper.ErrorPanic(err)
	workoutModel := models.Workout{
		IdAccount: workout.IdAccount,
		IdWorkout: workout.IdWorkout,
		Date:      workout.Date,
	}
	w.WorkoutsRepository.Save(workoutModel)
}

// Поиск всех тренировок
func (w *WorkoutsServiceImpl) FindAll() []responses.WorkoutsResponse {
	result := w.WorkoutsRepository.FindAll()

	var workouts []responses.WorkoutsResponse
	for _, value := range result {
		workout := responses.WorkoutsResponse{
			IdAccount: value.IdAccount,
			IdWorkout: value.IdWorkout,
			Date:      value.Date,
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
func (w *WorkoutsServiceImpl) FindById(workoutsId int) responses.WorkoutsResponse {
	workoutData, err := w.WorkoutsRepository.FindById(workoutsId)
	helper.ErrorPanic(err)

	workoutResponse := responses.WorkoutsResponse{
		IdAccount: workoutData.IdAccount,
		IdWorkout: workoutData.IdWorkout,
		Date:      workoutData.Date,
	}
	return workoutResponse
}

// Удаление тренировки
func (w *WorkoutsServiceImpl) Delete(workoutsId int) {
	w.WorkoutsRepository.Delete(workoutsId)
}
