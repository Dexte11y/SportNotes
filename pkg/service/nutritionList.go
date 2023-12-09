package service

import (
	"sportnotes/pkg/schemas"

	"sportnotes/pkg/repository"
)

type NutritionsListService struct {
	repo repository.NutritionList
}

func NewNutritionsListService(repo repository.NutritionList) *NutritionsListService {
	return &NutritionsListService{repo: repo}
}

func (s *NutritionsListService) CreateNutrition(idUser int, input schemas.Nutrition) (int, error) {
	return s.repo.CreateNutrition(idUser, input)
}

func (s *NutritionsListService) GetNutritionsByParam(id int, startpoint, endpoint string) ([]schemas.Nutrition, error) {
	return s.repo.GetNutritionsByParam(id, startpoint, endpoint)
}

func (s *NutritionsListService) DeleteNutrition(id int) error {
	return s.repo.DeleteNutrition(id)
}
