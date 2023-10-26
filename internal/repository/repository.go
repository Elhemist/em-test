package repository

import (
	"em-test/models"

	"github.com/jmoiron/sqlx"
)

type Person interface {
	AddPerson(models.PersonBD) (int, error)
	DeletePerson(int) error
}
type Repository struct {
	Person
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Person: NewPersonPostgres(db),
	}
}
