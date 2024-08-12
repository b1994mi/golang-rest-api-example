package user

import (
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
)

type request struct {
	ID          string    `json:"user_id" uri:"user_id" jwt:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Pin         string    `json:"pin"`
	CreatedAt   time.Time `json:"created_at"`
}

type handler struct {
	userRepo model.UserRepo
}

func NewHandler(
	userRepo model.UserRepo,
) *handler {
	return &handler{
		userRepo,
	}
}
