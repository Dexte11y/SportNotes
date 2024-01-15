package repository

import (
	"fmt"
	"strings"

	"sportnotes/internal/schemas"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ActivityListPostgres struct {
	db *sqlx.DB
}

func NewActivityListPostgres(db *sqlx.DB) *ActivityListPostgres {
	return &ActivityListPostgres{db: db}
}

func (r *ActivityListPostgres) CreateActivity(input schemas.Activity) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, approaches, repetitions, weight) VALUES ($1, $2, $3, $4) RETURNING  id ", activityTable)
	row := r.db.QueryRow(query, input.Name, input.Approaches, input.Repetitions, input.Weight)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ActivityListPostgres) GetAllActivity() ([]schemas.Activity, error) {
	var trainings []schemas.Activity

	query := fmt.Sprintf("SELECT id, id_workout ,name, approaches, repetitions, weight FROM %s", activityTable)
	err := r.db.Select(&trainings, query)

	return trainings, err
}

func (r *ActivityListPostgres) GetActivityByID(id int) (schemas.Activity, error) {
	var training schemas.Activity

	query := fmt.Sprintf("SELECT id, id_workout ,name, approaches, repetitions, weight FROM %s WHERE id = $1", activityTable)
	err := r.db.Get(&training, query, id)

	return training, err
}

func (r *ActivityListPostgres) UpdateActivity(id int, input schemas.UpdActivity) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("login=$%d", argID))
		args = append(args, *input.Name)
		argID++
	}

	if input.Approaches != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Approaches)
		argID++
	}

	if input.Repetitions != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argID))
		args = append(args, *input.Repetitions)
		argID++
	}

	if input.Weight != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argID))
		args = append(args, *input.Weight)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d", activityTable, setQuery, argID)
	args = append(args, id)

	logrus.Debug("updateQuery: ", query)
	logrus.Debug("args: ", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ActivityListPostgres) DeleteActivity(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", activityTable)
	_, err := r.db.Exec(query, id)

	return err
}
