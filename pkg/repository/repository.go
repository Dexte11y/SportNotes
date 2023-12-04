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

type Repository struct {
	// Authorisation
	FoodList
	NutritionList
	TrainingList
	UserList
	WorkoutList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		FoodList:      NewFoodsListPostgres(db),
		NutritionList: NewNutritionListPostgres(db),
		TrainingList:  NewTrainingsListPostgres(db),
		WorkoutList:   NewWorkoutsListPostgres(db),
		UserList:      NewUsersListPostgres(db),
		// Authorisation: NewAuthPostgres(db),
	}
}
