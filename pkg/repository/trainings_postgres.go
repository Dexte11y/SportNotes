package repository

import (
	"fmt"
	"strings"

	sportnotes "sportnotes"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TrainingsListPostgres struct {
	db *sqlx.DB
}

func NewTrainingsListPostgres(db *sqlx.DB) *TrainingsListPostgres {
	return &TrainingsListPostgres{db: db}
}

func (r *TrainingsListPostgres) CreateTraining(input sportnotes.Training) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (id, id_workout, name, approaches, repetitions, weight) VALUES ($1, $2, $3, $4, $5, $6) RETURNING  id ", trainingsTable)
	row := r.db.QueryRow(query, input.Id, input.IdWorkout, input.Name, input.Approaches, input.Repetitions, input.Weight)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TrainingsListPostgres) GetAllTrainings() ([]sportnotes.Training, error) {
	var trainings []sportnotes.Training

	query := fmt.Sprintf("SELECT id, id_workout ,name, approaches, repetitions, weight FROM %s", trainingsTable)
	err := r.db.Select(&trainings, query)

	return trainings, err
}

func (r *TrainingsListPostgres) GetTrainingById(id int) (sportnotes.Training, error) {
	var training sportnotes.Training

	query := fmt.Sprintf("SELECT id, id_workout ,name, approaches, repetitions, weight FROM %s WHERE id = $1", trainingsTable)
	err := r.db.Get(&training, query, id)

	return training, err
}

func (r *TrainingsListPostgres) UpdateTraining(id int, input sportnotes.UpdTraining) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("login=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Approaches != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Approaches)
		argId++
	}

	if input.Repetitions != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *input.Repetitions)
		argId++
	}

	if input.Weight != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Weight)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d", trainingsTable, setQuery, argId)
	args = append(args, id)

	logrus.Debug("updateQuery: ", query)
	logrus.Debug("args: ", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TrainingsListPostgres) DeleteTraining(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", trainingsTable)
	_, err := r.db.Exec(query, id)

	return err
}
