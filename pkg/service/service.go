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
	GetAllWorkouts() ([]sportnotes.Workout, error)
	GetWorkoutById(id int) (sportnotes.Workout, error)
	// UpdateWorkout(id int, input sportnotes.UpdWorkout) error
	DeleteWorkout(id int) error
}

type Service struct {
	// Authorisation
	TrainingList
	WorkoutList
	UserList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		TrainingList: NewTrainingsListService(repo.TrainingList),
		WorkoutList:  NewWorkoutsListService(repo.WorkoutList),
		UserList:     NewUsersListService(repo.UserList),

		// Authorisation: NewAuthService(repo.Authorisation),
	}
}
