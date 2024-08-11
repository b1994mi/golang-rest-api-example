# Golang REST API Example with BunRouter

## Summary

An example of a working back-end application with GORM as the DB ORM.

## Endpoints

| Endpoint      | HTTP   | Formdata Fields            | Body/Params                                                             |
| ------------- | ------ | ---------------------- | ---------------------------------------------------------------- |
| `/`    | GET    | Will give a simple response `{"message": "pong"}`.| |
| `/user/:id` | GET    | Get user data by Id   | |
| `/register`    | POST   | Create new user | Body (JSON): `first_name (string), last_name (string), phone_number (string), address (string), pin (string)`|
| `/login`    | POST   | Log in for user and returns jwt with refresh token as response | Body (JSON): `phone_number (string), pin (string)`|
| `/topup`    | POST   | Increase balance of user's wallet. Must use Auth Bearer JWT. | Body (JSON): `amount (number)`|
| `/transfer`    | POST   | Send money to another user. Must use Auth Bearer JWT. | Body (JSON): `target_user (string), amount (number), remarks (string)`|

# How to Start

You will need to install Go, RabbitMQ, and MySQL to run this app. You can easily run `docker-compose` if you have docker installed for easier local development setup. Assuming that you have installed Go, RabbitMQ and MySQL on your system, you can do these steps to start a local server:
- copy `.env-example` as `.env` and `dbconfig-example.yml` as `dbconfig.yml`
- fill those to files with the necessary values
- run migrations using `sql-migrate` command line tool:
    - install `sql-migrate` using `go install github.com/rubenv/sql-migrate/...@latest`
    - make sure your OS PATH is correctly setup, such as adding `export PATH=$PATH:/usr/local/go/bin` to `.profile` for linux
    - run `sel-migrate up` to mirate all the needed tables to your db
- run `go mod tidy` just to make sure all libraries are properly downloaded
- run the application using `go run .` or using VS Code debugger by opening the `main.go` file

# Decision Explanation

Some of the decisions that I made during development are as follows:
- I used `bunrouter` as it is one of the minimalist http library out there
- binding methods such as `ShouldBindJSON` are actually inspired (mostly copied) from Gin's methods
- `makeHandler` is inspired by Anthony GG's "Why Golang HTTP Handlers Should Return An Error" video and it wraps all the boilerplate that I usually encouter on golang projects using handler-usecase pattern; this abstracts the handler function away from many http implementation, thus hopefully makes it easier for dependency injection and testing
- `sql-migrate` as the DB migrate tool because I'm most familiar with that
- `rabbitmq` is used also because I'm already familiar with that
- `docker-compose` for easier local development setup