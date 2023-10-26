package service

import (
	repository "em-test/internal/repository"
)

type PersonService struct {
	repo repository.Person
}

func NewPersonService(repo repository.Person) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) AddPerson() error {
	// _, _ := repository.GetAge("Dima")
	// nya, _ := repository.GetNation("Dima")
	// nya, _ := repository.GetGender("Dima")
	return s.repo.AddPerson()
}
