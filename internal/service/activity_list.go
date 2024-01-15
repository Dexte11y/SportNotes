package service

import (
	"sportnotes/internal/schemas"

	"sportnotes/internal/repository"
)

type ActivityListService struct {
	repo repository.ActivityList
}

func NewActivityListService(repo repository.ActivityList) *ActivityListService {
	return &ActivityListService{repo: repo}
}

func (s *ActivityListService) CreateActivity(input schemas.Activity) (int, error) {
	return s.repo.CreateActivity(input)
}

func (s *ActivityListService) GetAllActivity() ([]schemas.Activity, error) {
	return s.repo.GetAllActivity()
}

func (s *ActivityListService) GetActivityByID(id int) (schemas.Activity, error) {
	return s.repo.GetActivityByID(id)
}

func (s *ActivityListService) UpdateActivity(id int, input schemas.UpdActivity) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateActivity(id, input)
}

func (s *ActivityListService) DeleteActivity(id int) error {
	return s.repo.DeleteActivity(id)
}
