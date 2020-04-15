# Gin API
CRUD API
Go + Gin + Gorm

## Requirement
* gin
* gorm
* json

## Usage

`sh reset_db.sh`

`go mod init`

`go build`

`go mod tidy`

`go run main.go`

## Operation

`curl http://localhost:8080/users -X GET`

`curl http://localhost:8080/users -X POST -H "Content-Type: application/json" -d '{"FirstName": "John", "LastName": "Doe"}'`

`curl http://localhost:8080/users/1 -X GET`

`curl http://localhost:8080/users/1 -X PUT -H "Content-Type: application/json" -d '{"FirstName": "Jane", "LastName": "Doe"}'`

`curl http://localhost:8080/users/1 -X DELETE`
