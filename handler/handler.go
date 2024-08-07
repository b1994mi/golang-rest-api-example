package handler

import (
	"github.com/b1994mi/golang-rest-api-example/model"
)

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
