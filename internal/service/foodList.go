package service

import (
	"sportnotes/internal/schemas"

	"sportnotes/internal/repository"
)

type FoodsListService struct {
	repo repository.FoodList
}

func NewFoodsListService(repo repository.FoodList) *FoodsListService {
	return &FoodsListService{repo: repo}
}

func (s *FoodsListService) CreateFood(input schemas.Food) (int, error) {
	return s.repo.CreateFood(input)
}

func (s *FoodsListService) GetAllFoods() ([]schemas.Food, error) {
	return s.repo.GetAllFoods()
}

func (s *FoodsListService) GetFoodByID(id int) (schemas.Food, error) {
	return s.repo.GetFoodByID(id)
}

// func (s *FoodsListService) UpdateFood(id int, input schemas.UpdFood) error {
// 	if err := input.Validate(); err != nil {
// 		return err
// 	}
// 	return s.repo.UpdateFood(id, input)
// }

func (s *FoodsListService) DeleteFood(id int) error {
	return s.repo.DeleteFood(id)
}
