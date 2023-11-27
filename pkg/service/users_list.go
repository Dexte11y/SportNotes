package service

import (
	sportnotes "sportnotes"

	"sportnotes/pkg/repository"
)

type UsersListService struct {
	repo repository.UserList
}

func NewUsersListService(repo repository.UserList) *UsersListService {
	return &UsersListService{repo: repo}
}

func (s *UsersListService) CreateUser(input sportnotes.User) (int, error) {
	return s.repo.CreateUser(input)
}

func (s *UsersListService) GetAllUsers() ([]sportnotes.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UsersListService) GetUserById(id int) (sportnotes.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UsersListService) UpdateUser(id int, input sportnotes.UpdUser) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateUser(id, input)
}

func (s *UsersListService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
