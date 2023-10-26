package service

import "em-test/internal/repository"

type Person interface {
	AddPerson() error
}
type Service struct {
	Person
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Person: NewPersonService(repos.Person),
	}
}
