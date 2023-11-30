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

type TrainingList interface {
	CreateTraining(training sportnotes.Training) (int, error)
	GetAllTrainings() ([]sportnotes.Training, error)
	GetTrainingById(id int) (sportnotes.Training, error)
	UpdateTraining(id int, input sportnotes.UpdTraining) error
	DeleteTraining(id int) error
}

type UserList interface {
	CreateUser(input sportnotes.User) (int, error)
	GetAllUsers() ([]sportnotes.User, error)
	GetUserById(id int) (sportnotes.User, error)
	UpdateUser(id int, input sportnotes.UpdUser) error
	DeleteUser(id int) error
}

type WorkoutList interface {
	CreateWorkout(workout sportnotes.Workout) (int, error)
	GetAllWorkouts() ([]sportnotes.WorkoutOutputAll, error)
	GetWorkoutById(id int) (sportnotes.WorkoutOutputById, error)
	// UpdateWorkout(id int, input sportnotes.UpdWorkout) error
	DeleteWorkout(id int) error
}

type Repository struct {
	// Authorisation
	TrainingList
	UserList
	WorkoutList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TrainingList: NewTrainingsListPostgres(db),
		WorkoutList:  NewWorkoutsListPostgres(db),
		UserList:     NewUsersListPostgres(db),
		// Authorisation: NewAuthPostgres(db),
	}
}
