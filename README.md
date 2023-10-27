# em-test

Решение тестового задания  для Effective Mobile на позицию Junior Go разработчик

# Инструкция по запуску

Для запуска необходимо создать postgres базу данных с именем nameTrace. Далее следует применить миграции. В случаее импользования утилиты migrate можно применить следующую команду.

```bash
migrate -path ./schema -database "postgres://postgres:qwerty@localhost:5432/nameTrace?sslmode=disable" up
```

После сервер пожно запустить с помощью следующей команды:

```bash
go run /cmd/emtest/main.go
```

# API

Данный сервис имеет четыре метода. Примеры запросов можно найти в файле EffectiveMobile.postman_collection.json

## Метод добавления человека в базу

### Post "/"

Данный метод принимает фио человека (отчество опционально) и добавляет в базу, обогащая возможными возрастом, полом и национальностью, используя запросы к сторонним api.

### Пример

```bash
curl --request POST \
  --url http://localhost:8080/person/ \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/2023.5.8' \
  --data '{
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich"
}'
```

## Метод удаления человека по id

### Delete "/"

Данный метод принимает id человека и удалет его из базы.

### Пример

```bash
curl --request DELETE \
  --url http://localhost:8080/person/ \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/2023.5.8' \
  --data '{
	"id": 6
}'
```

## Метод изменения сущности

### PATCH "/"

Данный метод принимает id человека и структуру включащуя в себя фио, возраст, пол и вес на которые требуется заменить данные пользоватедя по заданному id.

### Пример

```bash
curl --request PATCH \
  --url http://localhost:8080/person/ \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/2023.5.8' \
  --data '{
	"id": 8,
	"person": {
		"name": "Nykao",
		"surname": "Ushakov",
		"patronymic": "Vasilevich",
		"Age": 13,
		"Gender": "Helicopter",
		"Nationality": "KZ"
	}
}'
```

## Метод получения списка

### GET "/"

Данный метод принимает в себя размер и номер страницы для реализации пагинации, а также опциональные фио, возраст, пол и вес, в случае присутствия которых выбирает только людей соответствующих заданным параметрам.

### Пример

```bash
curl --request GET \
  --url http://localhost:8080/person/ \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Dmitriy",
	"surname": "Olegov",
	"patronymic": "Nikitich",
	"age": 40,
	"gender": "male",
	"nationality": "UA",
	"pageNumber": 2,
	"pageSize": 5
}'
```
