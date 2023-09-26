package repository

import (
	"fmt"
	"strings"

	server "github.com/deuuus/bmstu-rsoi"
	"github.com/jmoiron/sqlx"
)

type PersonPostgres struct {
	db *sqlx.DB
}

func NewPersonPostgres(db *sqlx.DB) *PersonPostgres {
	return &PersonPostgres{db: db}
}

func (r *PersonPostgres) CreatePerson(person server.Person) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, age, address, work) VALUES ($1, $2, $3, $4) RETURNING id", "persons")

	row := r.db.QueryRow(query, person.Name, person.Age, person.Address, person.Work)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PersonPostgres) GetPersonById(id int) (server.Person, error) {
	var person server.Person

	query := fmt.Sprintf("SELECT id, name, address, work, age FROM persons WHERE id = $1")

	if err := r.db.Get(&person, query, id); err != nil {
		return person, err
	}

	return person, nil
}

func (r *PersonPostgres) GetAllPersons() ([]server.Person, error) {
	var persons []server.Person

	query := fmt.Sprintf("SELECT id, name, address, work, age FROM persons")

	err := r.db.Select(&persons, query)

	return persons, err
}

func (r *PersonPostgres) DeletePersonById(id int) error {

	query := fmt.Sprintf("DELETE FROM persons WHERE id = $1")

	_, err := r.db.Exec(query, id)

	return err
}

func (r *PersonPostgres) UpdatePerson(id int, input server.PersonUpdate) (server.Person, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	var person server.Person

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}
	if input.Age != 0 {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, input.Age)
		argId++
	}
	if input.Address != "" {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, input.Address)
		argId++
	}
	if input.Work != "" {
		setValues = append(setValues, fmt.Sprintf("work=$%d", argId))
		args = append(args, input.Work)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE persons SET %s WHERE id = $%d", setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return person, err
	}

	selectQuery := "SELECT * FROM persons WHERE id = $1"
	if err := r.db.Get(&person, selectQuery, id); err != nil {
		return person, err
	}

	return person, nil
}
