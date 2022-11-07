# Goka Example Usage

## Summary

An example of a working back-end application with GORM as the DB ORM.

## Endpoints

| Endpoint      | HTTP   | Formdata Fields            | Body/Params                                                             |
| ------------- | ------ | ---------------------- | ---------------------------------------------------------------- |
| `/`    | GET    | Will give a simple response `{"message": "pong"}`.| |
| `/user` | GET    | Get user data by Id   | Body (JSON) : `id (number)`.|
| `/user`    | POST   | Create new user | Body (JSON): `email (string), name(string), phone_number(string), address (string), password (string), profile_image (string)`|
| `/verify`    | POST   | Verify a user | Body (JSON): `id (number)`|
