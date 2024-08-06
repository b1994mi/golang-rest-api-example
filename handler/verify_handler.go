package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/b1994mi/golang-rest-api-example/request"
	"github.com/uptrace/bunrouter"
)

func (h *handler) VerifyHandler(w http.ResponseWriter, req bunrouter.Request) error {
	body, err := io.ReadAll(req.Body)
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

	var m model.User

	err = h.db.Where("id", reqBody.ID).Find(&m).Error
	if err != nil {
		bunrouter.JSON(w, bunrouter.H{
			"message": err,
		})
	}

	m.IsUserActive = true
	now := time.Now()
	m.VerificationAt = &now

	err = h.db.Save(&m).Debug().Error
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
