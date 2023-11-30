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

func (s *WorkoutsListService) GetAllWorkouts() ([]sportnotes.WorkoutOutputAll, error) {
	return s.repo.GetAllWorkouts()
}

func (s *WorkoutsListService) GetWorkoutById(id int) (sportnotes.WorkoutOutputById, error) {
	return s.repo.GetWorkoutById(id)
}

// func (s *WorkoutsListService) UpdateWorkout(id int, input sportnotes.UpdWorkout) error {
// 	if err := input.Validate(); err != nil {
// 		return err
// 	}
// 	return s.repo.UpdateWorkout(id, input)
// }

func (s *WorkoutsListService) DeleteWorkout(id int) error {
	return s.repo.DeleteWorkout(id)
}
