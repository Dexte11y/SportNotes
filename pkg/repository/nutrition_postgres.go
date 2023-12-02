package repository

import (
	"fmt"
	"time"

	sportnotes "sportnotes"

	"github.com/jmoiron/sqlx"
)

type NutritionsListPostgres struct {
	db *sqlx.DB
}

func NewNutritionListPostgres(db *sqlx.DB) *NutritionsListPostgres {
	return &NutritionsListPostgres{db: db}
}

func (r *NutritionsListPostgres) CreateNutrition(input sportnotes.Nutrition) (int, error) {
	var id int
	currentData := time.Now().UTC()

	query := fmt.Sprintf("INSERT INTO %s (id, id_user, type, created_at) VALUES ($1, $2, $3, $4) RETURNING id", nutritionsTable)

	row := r.db.QueryRow(query, input.Id, input.IdUser, input.Type, currentData)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// func (r *NutritionsListPostgres) GetNutritionsByParam(id int) ([]sportnotes.NutritionOutputByParam, error) {
// var workoutsList []sportnotes.WorkoutOutputByParam
// var trainingsList []sportnotes.TrainingOutput
// currentDate := time.Now().UTC()

// queryWorkouts := fmt.Sprintf("SELECT id, id_user, type, created_at FROM %s WHERE id_user = $1 AND created_at >= $2 AND created_at <= $3", nutritionsTable)

// switch interval {
// case "all":
// 	previousDate := currentDate.AddDate(-100, 0, 0)
// 	r.db.Select(&workoutsList, queryWorkouts, id, previousDate, currentDate)
// case "year":
// 	previousDate := currentDate.AddDate(-1, 0, 0)
// 	r.db.Select(&workoutsList, queryWorkouts, id, previousDate, currentDate)
// case "month":
// 	previousDate := currentDate.AddDate(0, 0, -30)
// 	r.db.Select(&workoutsList, queryWorkouts, id, previousDate, currentDate)
// case "week":
// 	previousDate := currentDate.AddDate(0, 0, -7)
// 	r.db.Select(&workoutsList, queryWorkouts, id, previousDate, currentDate)
// }

// queryTrainings := fmt.Sprintf("SELECT id, id_workout, name, approaches, repetitions, weight FROM %s", trainingsTable)
// r.db.Select(&trainingsList, queryTrainings)

// mergetWorkouts := make([]sportnotes.WorkoutOutputByParam, 0)
// for _, valueWorkouts := range workoutsList {
// 	for _, valueTrainings := range trainingsList {
// 		if valueWorkouts.Id == valueTrainings.IdWorkout {
// 			valueWorkouts.TrainList = append(valueWorkouts.TrainList, valueTrainings)
// 		}
// 	}
// 	mergetWorkouts = append(mergetWorkouts, valueWorkouts)
// }

// return mergetWorkouts, nil
// }

// func (r *NutritionsListPostgres) UpdateNutrition(id int, input sportnotes.UpdWorkout) error {
// 	setValues := make([]string, 0)
// 	args := make([]interface{}, 0)
// 	argId := 1

// 	if input.CreatedAt != nil {
// 		setValues = append(setValues, fmt.Sprintf("login=$%d", argId))
// 		args = append(args, *input.CreatedAt)
// 		argId++
// 	}

// 	setQuery := strings.Join(setValues, ", ")

// 	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d", workoutsTable, setQuery, argId)
// 	args = append(args, id)

// 	logrus.Debug("updateQuery: ", query)
// 	logrus.Debug("args: ", args)

// 	_, err := r.db.Exec(query, args...)
// 	return err
// }

func (r *NutritionsListPostgres) DeleteNutrition(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", nutritionsTable)
	_, err := r.db.Exec(query, id)

	return err
}
