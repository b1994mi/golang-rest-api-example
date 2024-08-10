package user

import (
	"fmt"
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
	"golang.org/x/crypto/bcrypt"
)

func (h *handler) CreateHandler(req *request) (any, error) {
	now := time.Now()

	b := []byte(req.Pin)
	pinHashed, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to create pinHashed: %v", err)
	}

	tx := h.userRepo.StartTx()
	defer tx.Rollback()

	m, err := h.userRepo.Create(&model.User{
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Pin:         string(pinHashed),
		CreatedAt:   now,
	}, tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return m, nil
}
