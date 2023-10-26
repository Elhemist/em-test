package service

import (
	"em-test/internal/repository"
)

type PersonService struct {
	repo repository.Person
}

func NewPersonService(repo repository.Person) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) AddPerson() error {
	return s.repo.AddPerson()
}
