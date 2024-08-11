package transaction

import (
	"github.com/b1994mi/golang-rest-api-example/model"
)

type request struct {
	UserID string `jwt:"user_id"`

	Amount float64 `json:"amount"`
}

type handler struct {
	userRepo            model.UserRepo
	userTransactionRepo model.UserTransactionRepo
}

func NewHandler(
	userRepo model.UserRepo,
	userTransactionRepo model.UserTransactionRepo,
) *handler {
	return &handler{
		userRepo,
		userTransactionRepo,
	}
}
