package service

import (
	"sportnotes/pkg/schemas"

	"sportnotes/pkg/repository"
)

type UsersListService struct {
	repo repository.UserList
}

func NewUsersListService(repo repository.UserList) *UsersListService {
	return &UsersListService{repo: repo}
}

func (s *UsersListService) CreateUser(input schemas.User) (int, error) {
	return s.repo.CreateUser(input)
}

func (s *UsersListService) GetAllUsers() ([]schemas.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UsersListService) GetUserByID(id int) (schemas.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UsersListService) UpdateUser(id int, input schemas.UpdUser) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateUser(id, input)
}

func (s *UsersListService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
