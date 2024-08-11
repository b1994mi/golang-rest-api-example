package transaction

import (
	"github.com/b1994mi/golang-rest-api-example/message"
	"github.com/b1994mi/golang-rest-api-example/model"
)

type request struct {
	UserID string `jwt:"user_id"`

	Amount float64 `json:"amount"`

	TargetUser string `json:"target_user"`
	Remarks    string `json:"remarks"`
}

type handler struct {
	userRepo            model.UserRepo
	userTransactionRepo model.UserTransactionRepo
	transferRepo        message.TransferRepo
}

func NewHandler(
	userRepo model.UserRepo,
	userTransactionRepo model.UserTransactionRepo,
	transferRepo message.TransferRepo,
) *handler {
	return &handler{
		userRepo,
		userTransactionRepo,
		transferRepo,
	}
}
