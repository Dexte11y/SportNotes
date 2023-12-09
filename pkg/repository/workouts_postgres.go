package repository

import (
	"fmt"
	"sportnotes/pkg/schemas"
	"time"

	"github.com/jmoiron/sqlx"
)

type WorkoutsListPostgres struct {
	db *sqlx.DB
}

func NewWorkoutsListPostgres(db *sqlx.DB) *WorkoutsListPostgres {
	return &WorkoutsListPostgres{db: db}
}

func (r *WorkoutsListPostgres) CreateWorkout(idUser int, input schemas.Workout) (int, error) {
	var activityList []schemas.Activity
	var idWorkout int
	currentData := time.Now().UTC()

	activityList = append(activityList, input.ActivityList...)

	queryWorkouts := fmt.Sprintf("INSERT INTO %s (id_user, type, created_at) VALUES ($1, $2, $3) RETURNING id", workoutsTable)
	row := r.db.QueryRow(queryWorkouts, idUser, input.Type, currentData)
	if err := row.Scan(&idWorkout); err != nil {
		return 0, err
	}

	queryTrainings := fmt.Sprintf("INSERT INTO %s (id_workout, name, approaches, repetitions, weight) VALUES (%d, :name, :approaches, :repetitions, :weight)", activityTable, idWorkout)
	_, err := r.db.NamedExec(queryTrainings, activityList)
	if err != nil {
		return 0, err
	}

	return idWorkout, nil
}

func (r *WorkoutsListPostgres) GetWorkoutsByParam(id int, interval string) ([]schemas.Workout, error) {
	var workoutsList []schemas.Workout
	var trainingsList []schemas.Activity
	currentDate := time.Now().UTC()

	var intervalMap = map[string]time.Duration{"all": -876000 * time.Hour, "year": -8760 * time.Hour, "month": -720 * time.Hour, "week": -168 * time.Hour}
	previousDate := currentDate.Add(intervalMap[interval])

	queryWorkouts := fmt.Sprintf("SELECT id, id_user, type, created_at FROM %s WHERE id_user = $1 AND created_at >= $2 AND created_at <= $3", workoutsTable)
	err := r.db.Select(&workoutsList, queryWorkouts, id, previousDate, currentDate)
	if err != nil {
		return nil, err
	}

	queryTrainings := fmt.Sprintf("SELECT id, id_workout, name, approaches, repetitions, weight FROM %s", activityTable)
	err = r.db.Select(&trainingsList, queryTrainings)
	if err != nil {
		return nil, err
	}
	mergetWorkouts := make([]schemas.Workout, 0)
	for _, valueWorkouts := range workoutsList {
		for _, valueTrainings := range trainingsList {
			if valueWorkouts.ID == valueTrainings.IDWorkout {
				valueWorkouts.ActivityList = append(valueWorkouts.ActivityList, valueTrainings)
			}
		}
		mergetWorkouts = append(mergetWorkouts, valueWorkouts)
	}

	return mergetWorkouts, nil
}

// func (r *WorkoutsListPostgres) UpdateWorkout(id int, input schemas.UpdWorkout) error {
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

func (r *WorkoutsListPostgres) DeleteWorkout(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", workoutsTable)
	_, err := r.db.Exec(query, id)

	return err
}
