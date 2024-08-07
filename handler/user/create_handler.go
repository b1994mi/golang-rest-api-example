package user

import (
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/b1994mi/golang-rest-api-example/util"
	"github.com/uptrace/bunrouter"
)

func (h *handler) CreateHandler(r bunrouter.Request) (any, error) {
	var reqBody reqBody
	err := util.ShouldBindJSON(&reqBody, r)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	tx := h.userRepo.StartTx()
	defer tx.Rollback()

	m, err := h.userRepo.Create(&model.User{
		Email:        reqBody.Email,
		Name:         reqBody.Name,
		PhoneNumber:  reqBody.PhoneNumber,
		Address:      reqBody.Address,
		Password:     reqBody.Password,
		ProfileImage: reqBody.ProfileImage,
		CreatedAt:    now,
	}, tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return m, nil
}
