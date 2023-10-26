package service

import (
	"em-test/internal/repository"
	"em-test/models"
)

type Person interface {
	AddPerson(models.PersonInput) (int, error)
	DeletePerson(int) error
	EditPerson(models.PersonBD) error
	GetPersons(models.UserGetList) ([]models.PersonBD, error)
}
type Service struct {
	Person
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Person: NewPersonService(repos.Person),
	}
}
