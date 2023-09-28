package service

import (
	server "github.com/deuuus/bmstu-rsoi"
	"github.com/deuuus/bmstu-rsoi/pkg/repository"
)

type PersonService struct {
	repo repository.Person
}

func NewPersonService(repo repository.Person) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) CreatePerson(person server.Person) (int, error) {
	return s.repo.CreatePerson(person)
}

func (s *PersonService) GetPersonById(id int) (server.Person, error) {
	return s.repo.GetPersonById(id)
}

func (s *PersonService) GetAllPersons() ([]server.Person, error) {
	return s.repo.GetAllPersons()
}

func (s *PersonService) DeletePersonById(id int) (error) {
	return s.repo.DeletePersonById(id)
}

func (s *PersonService) UpdatePerson(id int, input server.PersonUpdate) (server.Person, error) {
	return s.repo.UpdatePerson(id, input)
}