package service

import (
	repository "em-test/internal/repository"
	"em-test/models"
	"fmt"
)

const (
	where = " WHERE "
	and   = "and "
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
func (s *PersonService) EditPerson(person models.PersonBD) error {
	return s.repo.EditPerson(person)
}
func (s *PersonService) GetPersons(person models.UserGetList) ([]models.PersonBD, error) {
	var settings = " WHERE "
	if person.Name != "" {
		settings += fmt.Sprintf("name = '%s' ", person.Name)
	}
	if person.Surname != "" {
		if settings != where {
			settings += and
		}
		settings += fmt.Sprintf("surname = '%s' ", person.Surname)
	}
	if person.Patronymic != "" {
		if settings != where {
			settings += and
		}
		settings += fmt.Sprintf("patronymic = '%s' ", person.Patronymic)
	}
	if person.Age != 0 {
		if settings != where {
			settings += and
		}
		settings += fmt.Sprintf("age = %d ", person.Age)
	}
	if person.Gender != "" {
		if settings != where {
			settings += and
		}
		settings += fmt.Sprintf("gender = '%s' ", person.Gender)
	}
	if person.Nationality != "" {
		if settings != where {
			settings += and
		}
		settings += fmt.Sprintf("nationality = '%s' ", person.Nationality)
	}
	if settings == where {
		settings = ""
	}
	if person.PageSize != 0 {
		pageNumber := person.PageNumber - 1
		if pageNumber < 0 {
			pageNumber = 0
		}
		settings += fmt.Sprintf("ORDER BY id OFFSET %d FETCH NEXT %d ROWS ONLY", pageNumber*person.PageSize, person.PageSize)
	}
	return s.repo.GetPersons(settings)
}
