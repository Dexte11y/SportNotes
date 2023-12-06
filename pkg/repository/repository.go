package repository

import (
	"sportnotes/pkg/schemas"

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
	CreateFood(food schemas.Food) (int, error)
	GetAllFoods() ([]schemas.Food, error)
	GetFoodById(id int) (schemas.Food, error)
	// UpdateFood(id int, input sportnotes.UpdFood) error
	DeleteFood(id int) error
}

type NutritionList interface {
	CreateNutrition(nutrition schemas.Nutrition) (int, error)
	GetNutritionsByParam(id int, startpoint, endpoint string) ([]schemas.Nutrition, error)
	// UpdateNutrition(id int, input sportnotes.UpdNutrition) error
	DeleteNutrition(id int) error
}

type TrainingList interface {
	CreateTraining(training schemas.Training) (int, error)
	GetAllTrainings() ([]schemas.Training, error)
	GetTrainingById(id int) (schemas.Training, error)
	UpdateTraining(id int, input schemas.UpdTraining) error
	DeleteTraining(id int) error
}

type UserList interface {
	CreateUser(user schemas.User) (int, error)
	GetAllUsers() ([]schemas.User, error)
	GetUserById(id int) (schemas.User, error)
	UpdateUser(id int, input schemas.UpdUser) error
	DeleteUser(id int) error
}

type WorkoutList interface {
	CreateWorkout(idUser int, workout schemas.Workout) (int, error)
	GetWorkoutsByParam(id int, input string) ([]schemas.Workout, error)
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
