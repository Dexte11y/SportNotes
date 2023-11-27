package repository

import (
	"fmt"
	"strings"

	sportnotes "sportnotes"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type WorkoutsListPostgres struct {
	db *sqlx.DB
}

func NewWorkoutsListPostgres(db *sqlx.DB) *WorkoutsListPostgres {
	return &WorkoutsListPostgres{db: db}
}

func (r *WorkoutsListPostgres) CreateWorkout(input sportnotes.Workout) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (id, id_user, created_at) VALUES ($1, $2, $3) RETURNING id",
		workoutsTable)

	row := r.db.QueryRow(query, input.Id, input.IdUser, input.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *WorkoutsListPostgres) GetAllWorkouts() ([]sportnotes.Workout, error) {
	var workouts []sportnotes.Workout

	query := fmt.Sprintf("SELECT id, id_user, created_at FROM %s", workoutsTable)
	err := r.db.Select(&workouts, query)

	return workouts, err
}

func (r *WorkoutsListPostgres) GetWorkoutById(id int) (sportnotes.Workout, error) {
	var workout sportnotes.Workout

	query := fmt.Sprintf("SELECT id, id_user, created_at FROM %s WHERE id = $1", workoutsTable)
	err := r.db.Get(&workout, query, id)

	return workout, err
}

func (r *WorkoutsListPostgres) UpdateWorkout(id int, input sportnotes.UpdWorkout) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Login != nil {
		setValues = append(setValues, fmt.Sprintf("login=$%d", argId))
		args = append(args, *input.Login)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d", workoutsTable, setQuery, argId)
	args = append(args, id)

	logrus.Debug("updateQuery: ", query)
	logrus.Debug("args: ", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *WorkoutsListPostgres) DeleteWorkout(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", workoutsTable)
	_, err := r.db.Exec(query, id)

	return err
}