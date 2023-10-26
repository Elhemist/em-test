package repository

import (
	"github.com/jmoiron/sqlx"
)

type Person interface {
	AddPerson() error
}
type Repository struct {
	Person
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Person: NewPersonPostgres(db),
	}
}
