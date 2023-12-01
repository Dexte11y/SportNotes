package repository

import (
	"fmt"
	"time"

	sportnotes "sportnotes"

	"github.com/jmoiron/sqlx"
)

type WorkoutsListPostgres struct {
	db *sqlx.DB
}

func NewWorkoutsListPostgres(db *sqlx.DB) *WorkoutsListPostgres {
	return &WorkoutsListPostgres{db: db}
}

func (r *WorkoutsListPostgres) CreateWorkout(input sportnotes.Workout) (int, error) {
	var id int
	data := time.Now().Format("2006-01-02")
	query := fmt.Sprintf("INSERT INTO %s (id, id_user, type, created_at) VALUES ($1, $2, $3, $4) RETURNING id", workoutsTable)

	row := r.db.QueryRow(query, input.Id, input.IdUser, input.Type, data)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *WorkoutsListPostgres) GetAllWorkouts() ([]sportnotes.WorkoutOutputAll, error) {
	var workoutsList []sportnotes.WorkoutOutputAll
	var trainList []sportnotes.TrainingOutput

	queryWorkouts := fmt.Sprintf("SELECT id, type, created_at FROM %s", workoutsTable)
	err := r.db.Select(&workoutsList, queryWorkouts)

	queryTrainings := fmt.Sprintf("SELECT id, id_workout, name, approaches, repetitions, weight FROM %s", trainingsTable)
	r.db.Select(&trainList, queryTrainings)

	mergetWorkouts := make([]sportnotes.WorkoutOutputAll, 0)
	for _, valueWorkouts := range workoutsList {
		for _, valueTrainings := range trainList {
			if valueWorkouts.Id == valueTrainings.IdWorkout {
				valueWorkouts.TrainList = append(valueWorkouts.TrainList, valueTrainings)
			}
		}
		mergetWorkouts = append(mergetWorkouts, valueWorkouts)
	}
	return mergetWorkouts, err
}

func (r *WorkoutsListPostgres) GetWorkoutById(id int) (sportnotes.WorkoutOutputById, error) {
	var workoutsList sportnotes.WorkoutOutputById
	var trainList []sportnotes.TrainingOutput

	queryWorkouts := fmt.Sprintf("SELECT id, type, created_at FROM %s WHERE id = $1", workoutsTable)
	err := r.db.Get(&workoutsList, queryWorkouts, id)

	queryTrainings := fmt.Sprintf("SELECT id, id_workout, name, approaches, repetitions, weight FROM %s WHERE id_workout = $1", trainingsTable)
	r.db.Select(&trainList, queryTrainings, id)

	workoutsList.TrainList = trainList

	return workoutsList, err
}

// func (r *WorkoutsListPostgres) UpdateWorkout(id int, input sportnotes.UpdWorkout) error {
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
