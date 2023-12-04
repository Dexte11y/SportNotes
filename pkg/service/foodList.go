package service

import (
	sportnotes "sportnotes"

	"sportnotes/pkg/repository"
)

type FoodsListService struct {
	repo repository.FoodList
}

func NewFoodsListService(repo repository.FoodList) *FoodsListService {
	return &FoodsListService{repo: repo}
}

func (s *FoodsListService) CreateFood(input sportnotes.Food) (int, error) {
	return s.repo.CreateFood(input)
}

func (s *FoodsListService) GetAllFoods() ([]sportnotes.Food, error) {
	return s.repo.GetAllFoods()
}

func (s *FoodsListService) GetFoodById(id int) (sportnotes.Food, error) {
	return s.repo.GetFoodById(id)
}

// func (s *FoodsListService) UpdateFood(id int, input sportnotes.UpdFood) error {
// 	if err := input.Validate(); err != nil {
// 		return err
// 	}
// 	return s.repo.UpdateFood(id, input)
// }

func (s *FoodsListService) DeleteFood(id int) error {
	return s.repo.DeleteFood(id)
}
