package handler

import (
	"encoding/json"
	"io"
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/b1994mi/golang-rest-api-example/request"
	"github.com/uptrace/bunrouter"
)

func (h *handler) VerifyHandler(r bunrouter.Request) (any, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var reqBody request.User
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		return nil, err
	}

	var m model.User

	err = h.db.Where("id", reqBody.ID).Find(&m).Error
	if err != nil {
		return nil, err
	}

	m.IsUserActive = true
	now := time.Now()
	m.VerificationAt = &now

	err = h.db.Save(&m).Debug().Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
