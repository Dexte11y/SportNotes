package repository

import (
	"fmt"

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
	query := fmt.Sprintf("INSERT INTO %s (id, id_user, type, created_at) VALUES ($1, $2, $3, $4) RETURNING id", workoutsTable)

	row := r.db.QueryRow(query, input.Id, input.IdUser, input.Type, input.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *WorkoutsListPostgres) GetAllWorkouts() ([]sportnotes.Workout, error) {
	var workouts []sportnotes.Workout

	query := fmt.Sprintf("SELECT id, id_user, type, created_at FROM %s", workoutsTable)
	err := r.db.Select(&workouts, query)

	return workouts, err
}

func (r *WorkoutsListPostgres) GetWorkoutById(id int) (sportnotes.WorkoutOutputById, error) {
	var workout sportnotes.WorkoutOutputById
	var trainList []sportnotes.Train

	query := fmt.Sprintf("SELECT id, type, created_at FROM %s WHERE id = $1", workoutsTable)
	err := r.db.Get(&workout, query, id)

	stmt := fmt.Sprintf("SELECT id, name, approaches, repetitions, weight FROM %s WHERE id_workout = $1", trainingsTable)
	r.db.Select(&trainList, stmt, id)

	workout.TrainList = trainList

	return workout, err
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
