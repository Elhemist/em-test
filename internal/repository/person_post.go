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

// func (r *SegmentPostgres) CreateSegment(name string) (int, error) {

// 	var userSeg segment.Segment
// 	if name != "" {
// 		query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1)", segmentTable)
// 		_, err := r.db.Exec(query, name)
// 		if err != nil {
// 			return userSeg.Id, fmt.Errorf("Segment insert error")
// 		}
// 		query = fmt.Sprintf("SELECT id, name FROM %s WHERE name=$1 ", segmentTable)
// 		err = r.db.Get(&userSeg, query, name)
// 		return userSeg.Id, err
// 	} else {
// 		return userSeg.Id, fmt.Errorf("Empty name")
// 	}
// }
