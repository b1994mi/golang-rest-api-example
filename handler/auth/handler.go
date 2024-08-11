package auth

import (
	"github.com/b1994mi/golang-rest-api-example/model"
)

type request struct {
	RefreshToken string `json:"refresh_token"`
	PhoneNumber  string `json:"phone_number"`
	PIN          string `json:"pin"`
}

type handler struct {
	userRepo      model.UserRepo
	userTokenRepo model.UserTokenRepo
}

func NewHandler(
	userRepo model.UserRepo,
	userTokenRepo model.UserTokenRepo,
) *handler {
	return &handler{
		userRepo,
		userTokenRepo,
	}
}
