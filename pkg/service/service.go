package service

import (
	sportnotes "sportnotes"

	"sportnotes/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

// type Authorisation interface {
// 	CreateDoctor(doctor sportnotes.Doctor) (int, error)
// 	GenerateToken(login, password string) (string, error)
// 	ParseToken(token string) (int, error)
// }

// type DoctorList interface {
// 	GetAll() ([]sportnotes.Doctor, error)
// 	GetById(id int) (sportnotes.Doctor, error)
// }

type UserList interface {
	CreateUser(input sportnotes.User) (int, error)
	GetAll() ([]sportnotes.User, error)
	GetById(id int) (sportnotes.User, error)
	UpdateUser(id int, input sportnotes.UpdUser) error
	DeleteUser(id int) error
}

// type VisitList interface {
// 	CreateVisit(input sportnotes.Visit) (int, error)
// 	GetAll() ([]sportnotes.VisitOutput, error)
// 	GetById(id int) (sportnotes.VisitOutput, error)
// 	UpdateVisit(id int, input sportnotes.UpdVisit) error
// 	DeleteVisit(id int) error
// }

type Service struct {
	// Authorisation
	// DoctorList
	UserList
	// VisitList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserList: NewUsersListService(repo.UserList),

		// Authorisation: NewAuthService(repo.Authorisation),
		// DoctorList:    NewDoctorsListService(repo.DoctorList),
		// VisitList:     NewVisitsListPostgres(repo.VisitList),
	}
}
