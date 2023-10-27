up:
	go run /cmd/emtest/main.go
migrations-up:
	migrate -path ./schema -database "postgres://postgres:qwerty@localhost:5432/nameTrace?sslmode=disable" up
migrations-down:
	migrate -path ./schema -database "postgres://postgres:qwerty@localhost:5432/nameTrace?sslmode=disable" down
