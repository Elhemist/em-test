package models

type PersonBD struct {
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         int    `json:"age,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
}
type PersonInput struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}



{
	"id": 5,
	"person":
	{
"name": "Dmitriy",
"surname": "Ushakov",
"patronymic": "Vasilevich",
	"Age": 13,
		"Gender": "Helicopter",
"Nationality": "KZ"
}
}