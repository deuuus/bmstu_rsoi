package service

import (
	server "github.com/deuuus/bmstu-rsoi"
	"github.com/deuuus/bmstu-rsoi/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Person interface {
	CreatePerson(person server.Person) (int, error)
	GetPersonById(id int) (server.Person, error)
	GetAllPersons() ([]server.Person, error)
	DeletePersonById(id int) (error)
	UpdatePerson(id int, input server.PersonUpdate) (server.Person, error)
}

type Service struct {
	Person
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Person: NewPersonService(repos.Person)}
}
