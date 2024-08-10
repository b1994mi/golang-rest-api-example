package user

import (
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
)

func (h *handler) CreateHandler(req *request) (any, error) {
	now := time.Now()

	tx := h.userRepo.StartTx()
	defer tx.Rollback()

	m, err := h.userRepo.Create(&model.User{
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Pin:         req.Pin,
		CreatedAt:   now,
	}, tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return m, nil
}
