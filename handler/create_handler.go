package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/b1994mi/golang-rest-api-example/request"
	"github.com/uptrace/bunrouter"
)

func (h *handler) CreateHandler(w http.ResponseWriter, req bunrouter.Request) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		bunrouter.JSON(w, bunrouter.H{"message": err})
		return nil
	}

	var reqBody request.User
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		bunrouter.JSON(w, bunrouter.H{"message": err})
		return nil
	}
	now := time.Now()
	err = h.db.Create(&model.User{
		Email:          "",
		Name:           reqBody.Name,
		PhoneNumber:    "",
		Address:        "",
		Password:       "",
		IsUserActive:   false,
		VerificationAt: now,
		ProfileImage:   "",
		CreatedAt:      now,
	}).Error
	if err != nil {
		bunrouter.JSON(w, bunrouter.H{
			"message": err,
		})
	}

	bunrouter.JSON(w, bunrouter.H{
		"acknowledge": true,
	})
	return nil
}
