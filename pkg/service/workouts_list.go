package service

import (
	sportnotes "sportnotes"

	"sportnotes/pkg/repository"
)

type WorkoutsListService struct {
	repo repository.WorkoutList
}

func NewWorkoutsListService(repo repository.WorkoutList) *WorkoutsListService {
	return &WorkoutsListService{repo: repo}
}

func (s *WorkoutsListService) CreateWorkout(input sportnotes.Workout) (int, error) {
	return s.repo.CreateWorkout(input)
}

func (s *WorkoutsListService) GetWorkoutsByParam(id int, input string) ([]sportnotes.WorkoutOutputByParam, error) {
	return s.repo.GetWorkoutsByParam(id, input)
}

func (s *WorkoutsListService) GetWorkoutById(id int) (sportnotes.WorkoutOutputById, error) {
	return s.repo.GetWorkoutById(id)
}

func (s *WorkoutsListService) DeleteWorkout(id int) error {
	return s.repo.DeleteWorkout(id)
}
