package handler

import (
	"encoding/json"
	"io"
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/b1994mi/golang-rest-api-example/request"
	"github.com/uptrace/bunrouter"
)

func (h *handler) CreateHandler(r bunrouter.Request) (any, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var reqBody request.User
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	m := model.User{
		Email:        reqBody.Email,
		Name:         reqBody.Name,
		PhoneNumber:  reqBody.PhoneNumber,
		Address:      reqBody.Address,
		Password:     reqBody.Password,
		ProfileImage: reqBody.ProfileImage,
		CreatedAt:    now,
	}

	err = h.db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
