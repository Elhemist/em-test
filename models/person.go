package models

type PersonBD struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name,omitempty"  db:"name"`
	Surname     string `json:"surname,omitempty"  db:"surname"`
	Patronymic  string `json:"patronymic,omitempty"  db:"patronymic"`
	Age         int    `json:"age,omitempty"  db:"age"`
	Gender      string `json:"gender,omitempty"  db:"gender"`
	Nationality string `json:"nationality,omitempty"  db:"nationality"`
}
type PersonInput struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

type UserGetList struct {
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         int    `json:"age,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	PageNumber  int    `json:"pageNumber"`
	PageSize    int    `json:"pageSize"`
}
