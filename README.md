# Go example projects

[![Go Reference](https://pkg.go.dev/badge)

This repository contains a collection of Go programs and libraries that
demonstrate the language, standard libraries, and tools.

## Clone the project

```
$ git clone https://github.com/regasyahfika/go_restapi.git
$ cd backer
$ touch .env

// file .env
DB_USERNAME=your_user
DB_PASSWORD=your_password
DATABASE_HOST=127.0.0.1:3306
DB_DATABASE=your_database

note : create your database and table auto migrate
```

## How to Run

```
$ go run main.go
```

## Testing Postman

```
Get all data : http://localhost:8080/v1/book 
Get detail data : http://localhost:8080/v1/book/1
Create data : http://localhost:8080/v1/book
Update data : http://localhost:8080/v1/book/1
Delete data : http://localhost:8080/v1/book/1
```

## Package and Framework
* github.com/gin-gonic/gin
* github.com/go-playground/validator/v10
* github.com/joho/godotenv
* gorm.io/driver/mysql
* gorm.io/gorm