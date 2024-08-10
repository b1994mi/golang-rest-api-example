package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/b1994mi/golang-rest-api-example/util"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (h *handler) CreateHandler(req *request) (any, error) {
	now := time.Now()

	user, err := h.userRepo.FindOneBy(map[string]any{
		"phone_number": req.PhoneNumber,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user != nil {
		return nil, util.New409Res("phone number already registered")
	}

	b := []byte(req.Pin)
	pinHashed, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to create %v as pinHashed: %v", req.Pin, err)
	}

	tx := h.userRepo.StartTx()
	defer tx.Rollback()

	m, err := h.userRepo.Create(&model.User{
		ID:          uuid.New().String(),
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
