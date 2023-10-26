package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	personsTable = "persons"
)

type Config struct {
	Host     string
	Port     string
	Usename  string
	Password string
	DBName   string
	SSLmode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Usename, cfg.DBName, cfg.Password, cfg.SSLmode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitTables(db *sqlx.DB) {
	query := `
    DROP TABLE IF EXISTS persons;
	CREATE TABLE persons(
		id          serial          primary key,
		name        varchar(255)    not null,
		surname     varchar(255)    not null,
		patronymic  varchar(255),
		age         numeric         not null, 
		gender      varchar(255)    not null,
		nationality varchar(255)    not null,
	);
	INSERT INTO persons(name, surname, age, gender, nationality) VALUES ('1', '1', 11, '1', '1');
	`
	_, _ = db.Exec(query)
}
