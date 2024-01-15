package service

import (
	"sportnotes/internal/schemas"

	"sportnotes/internal/repository"
)

type WorkoutsListService struct {
	repo repository.WorkoutList
}

func NewWorkoutsListService(repo repository.WorkoutList) *WorkoutsListService {
	return &WorkoutsListService{repo: repo}
}

func (s *WorkoutsListService) CreateWorkout(idUser int, input schemas.Workout) (int, error) {
	return s.repo.CreateWorkout(idUser, input)
}

func (s *WorkoutsListService) GetWorkoutsByParam(id int, input string) ([]schemas.Workout, error) {
	return s.repo.GetWorkoutsByParam(id, input)
}

func (s *WorkoutsListService) DeleteWorkout(id int) error {
	return s.repo.DeleteWorkout(id)
}
