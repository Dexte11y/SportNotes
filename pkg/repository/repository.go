package repository

import (
	sportnotes "sportnotes"

	"github.com/jmoiron/sqlx"
)

// type Authorisation interface {
// 	CreateDoctor(doctor medapp.Doctor) (int, error)
// 	GetDoctor(login, password string) (medapp.Doctor, error)
// }

// type DoctorList interface {
// 	GetAll() ([]medapp.Doctor, error)
// 	GetById(id int) (medapp.Doctor, error)
// }

type UserList interface {
	CreateUser(input sportnotes.User) (int, error)
	GetAll() ([]sportnotes.User, error)
	GetById(id int) (sportnotes.User, error)
	UpdateUser(id int, input sportnotes.UpdUser) error
	DeleteUser(id int) error
}

// type VisitList interface {
// 	CreateVisit(input medapp.Visit) (int, error)
// 	GetAll() ([]medapp.VisitOutput, error)
// 	GetById(id int) (medapp.VisitOutput, error)
// 	UpdateVisit(id int, input medapp.UpdVisit) error
// 	DeleteVisit(id int) error
// }

type Repository struct {
	// Authorisation
	// DoctorList
	UserList
	// VisitList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserList: NewUsersListPostgres(db),
		// Authorisation: NewAuthPostgres(db),
		// DoctorList:    NewDoctorsListPostgres(db),
		// VisitList:     NewVisitsListPostgres(db),
	}
}
