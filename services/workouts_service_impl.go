package services

import (
	"SportNotes/data/requests"
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

func (w *WorkoutsServiceImpl) Create(workout requests.CreateWorkoutsRequest) {
	err := w.Validate.Struct(workout)
	helper.ErrorPanic(err)
	workoutModel := models.Workouts{
		Id:   workout.Id,
		Name: workout.Name,
	}
	w.WorkoutsRepository.Save(workoutModel)
}
