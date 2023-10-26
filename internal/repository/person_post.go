package repository

import (
	"em-test/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PersonPostgres struct {
	db *sqlx.DB
}

func NewPersonPostgres(db *sqlx.DB) *PersonPostgres {
	return &PersonPostgres{db: db}
}

func (r *PersonPostgres) AddPerson(person models.PersonBD) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (Name, Surname, Patronymic, Age, Gender, Nationality) VALUES ($1, $2, $3, $4, $5, $6)", personsTable)
	_, err := r.db.Exec(query, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality)
	if err != nil {
		return -1, fmt.Errorf("Person insert error")
	}
	query = fmt.Sprintf("SELECT id FROM %s WHERE Name=$1 and Surname=$2 and Patronymic=$3", personsTable)
	err = r.db.Get(&id, query, person.Name, person.Surname, person.Patronymic)
	if err != nil {
		return -1, fmt.Errorf("Person select error")
	}
	return int(id), nil
}

func (r *PersonPostgres) DeletePerson(id int) error {
	query := fmt.Sprintf("DELETE FROM  %s WHERE id =$1", personsTable)
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Row with person id %v delete error", id)
	}
	return err
}
func (r *PersonPostgres) EditPerson(id int, person models.PersonBD) error {
	query := fmt.Sprintf("UPDATE %s SET Name=$1, Surname=$2, Patronymic=$3, Age=$4, Gender=$5, Nationality=$6 WHERE id =$7", personsTable)
	res, err := r.db.Exec(query, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Row with person id %v update error", id)
	}
	return err
}
