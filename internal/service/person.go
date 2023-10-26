package service

import (
	repository "em-test/internal/repository"
	"em-test/models"
)

type PersonService struct {
	repo repository.Person
}

func NewPersonService(repo repository.Person) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) AddPerson(personIn models.PersonInput) (int, error) {
	var err error
	person := models.PersonBD{
		Name:       personIn.Name,
		Surname:    personIn.Surname,
		Patronymic: personIn.Patronymic,
	}
	person.Age, err = repository.GetAge(personIn.Name)
	if err != nil {
		return -1, err
	}
	person.Nationality, err = repository.GetNation(personIn.Name)
	if err != nil {
		return -1, err
	}
	person.Gender, err = repository.GetGender(personIn.Name)
	if err != nil {
		return -1, err
	}
	return s.repo.AddPerson(person)
}

func (s *PersonService) DeletePerson(id int) error {
	return s.repo.DeletePerson(id)
}
