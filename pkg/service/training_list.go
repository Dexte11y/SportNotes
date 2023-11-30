package service

import (
	sportnotes "sportnotes"

	"sportnotes/pkg/repository"
)

type TrainingsListService struct {
	repo repository.TrainingList
}

func NewTrainingsListService(repo repository.TrainingList) *TrainingsListService {
	return &TrainingsListService{repo: repo}
}

func (s *TrainingsListService) CreateTraining(input sportnotes.Training) (int, error) {
	return s.repo.CreateTraining(input)
}

func (s *TrainingsListService) GetAllTrainings() ([]sportnotes.Training, error) {
	return s.repo.GetAllTrainings()
}

func (s *TrainingsListService) GetTrainingById(id int) (sportnotes.Training, error) {
	return s.repo.GetTrainingById(id)
}

func (s *TrainingsListService) UpdateTraining(id int, input sportnotes.UpdTraining) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateTraining(id, input)
}

func (s *TrainingsListService) DeleteTraining(id int) error {
	return s.repo.DeleteTraining(id)
}
