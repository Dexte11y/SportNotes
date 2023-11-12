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
	workoutModel := models.Workouts{
		Id:   workout.Id,
		Name: workout.Name,
	}
	w.WorkoutsRepository.Save(workoutModel)
}

// Поиск всех тренировок
func (w *WorkoutsServiceImpl) FindAll() []responses.WorkoutsResponse {
	result := w.WorkoutsRepository.FindAll()

	var workouts []responses.WorkoutsResponse
	for _, value := range result {
		workout := responses.WorkoutsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		workouts = append(workouts, workout)
	}

	return workouts
}

// Обновление тренировки
func (w *WorkoutsServiceImpl) Update(workouts requests.UpdateWorkoutsRequest) {
	workoutData, err := w.WorkoutsRepository.FindById(workouts.Id)
	helper.ErrorPanic(err)
	workoutData.Name = workouts.Name
	w.WorkoutsRepository.Update(workoutData)
}

// Поиск тренировки по Id
func (w *WorkoutsServiceImpl) FindById(workoutsId int) responses.WorkoutsResponse {
	workoutData, err := w.WorkoutsRepository.FindById(workoutsId)
	helper.ErrorPanic(err)

	workoutResponse := responses.WorkoutsResponse{
		Id:   workoutData.Id,
		Name: workoutData.Name,
	}
	return workoutResponse
}
