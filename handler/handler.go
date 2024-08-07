package handler

import (
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
)

type reqBody struct {
	ID             int        `json:"id"`
	Email          string     `json:"email"`
	Name           string     `json:"name"`
	PhoneNumber    string     `json:"phone_number"`
	Address        string     `json:"address"`
	Password       string     `json:"password"`
	IsUserActive   bool       `json:"is_user_active"`
	VerificationAt *time.Time `json:"verification_at"`
	ProfileImage   string     `json:"profile_image"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
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
