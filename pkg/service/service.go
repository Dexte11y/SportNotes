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

type FoodList interface {
	CreateFood(food sportnotes.Food) (int, error)
	GetAllFoods() ([]sportnotes.Food, error)
	GetFoodById(id int) (sportnotes.Food, error)
	// UpdateFood(id int, input sportnotes.UpdFood) error
	DeleteFood(id int) error
}

type NutritionList interface {
	CreateNutrition(nutrition sportnotes.Nutrition) (int, error)
	GetNutritionsByParam(id int, startpoint, endpoint string) ([]sportnotes.NutritionOutputByParam, error)
	// UpdateNutrition(id int, input sportnotes.UpdNutrition) error
	DeleteNutrition(id int) error
}

type TrainingList interface {
	CreateTraining(training sportnotes.Training) (int, error)
	GetAllTrainings() ([]sportnotes.Training, error)
	GetTrainingById(id int) (sportnotes.Training, error)
	UpdateTraining(id int, input sportnotes.UpdTraining) error
	DeleteTraining(id int) error
}

type UserList interface {
	CreateUser(user sportnotes.User) (int, error)
	GetAllUsers() ([]sportnotes.User, error)
	GetUserById(id int) (sportnotes.User, error)
	UpdateUser(id int, input sportnotes.UpdUser) error
	DeleteUser(id int) error
}

type WorkoutList interface {
	CreateWorkout(workout sportnotes.Workout) (int, error)
	GetWorkoutsByParam(id int, input string) ([]sportnotes.WorkoutOutputByParam, error)
	// UpdateWorkout(id int, input sportnotes.UpdWorkout) error
	DeleteWorkout(id int) error
}

type Service struct {
	// Authorisation
	FoodList
	NutritionList
	TrainingList
	WorkoutList
	UserList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		FoodList:      NewFoodsListService(repo.FoodList),
		NutritionList: NewNutritionsListService(repo.NutritionList),
		TrainingList:  NewTrainingsListService(repo.TrainingList),
		WorkoutList:   NewWorkoutsListService(repo.WorkoutList),
		UserList:      NewUsersListService(repo.UserList),

		// Authorisation: NewAuthService(repo.Authorisation),
	}
}
