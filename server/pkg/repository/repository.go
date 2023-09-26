package repository

import (
	server "github.com/deuuus/bmstu-rsoi"
	"github.com/jmoiron/sqlx"
)

type Person interface {
	CreatePerson(person server.Person) (int, error)
	GetPersonById(id int) (server.Person, error)
	GetAllPersons() ([]server.Person, error)
	DeletePersonById(id int) (error)
	UpdatePerson(id int, input server.PersonUpdate) (server.Person, error)
}

type Repository struct {
	Person
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Person: NewPersonPostgres(db),
	}
}
