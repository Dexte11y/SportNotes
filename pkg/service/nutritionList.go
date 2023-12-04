package service

import (
	sportnotes "sportnotes"

	"sportnotes/pkg/repository"
)

type NutritionsListService struct {
	repo repository.NutritionList
}

func NewNutritionsListService(repo repository.NutritionList) *NutritionsListService {
	return &NutritionsListService{repo: repo}
}

func (s *NutritionsListService) CreateNutrition(input sportnotes.Nutrition) (int, error) {
	return s.repo.CreateNutrition(input)
}

func (s *NutritionsListService) GetNutritionsByParam(id int, startpoint, endpoint string) ([]sportnotes.NutritionOutputByParam, error) {
	return s.repo.GetNutritionsByParam(id, startpoint, endpoint)
}

func (s *NutritionsListService) DeleteNutrition(id int) error {
	return s.repo.DeleteNutrition(id)
}
