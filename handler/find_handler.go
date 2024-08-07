package handler

import (
	"encoding/json"
	"io"

	"github.com/b1994mi/golang-rest-api-example/request"
	"github.com/uptrace/bunrouter"
)

func (h *handler) FindHandler(r bunrouter.Request) (any, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var reqBody request.User
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		return nil, err
	}

	m, err := h.userRepo.FindOneBy(map[string]interface{}{
		"id": reqBody.ID,
	})
	if err != nil {
		return nil, err
	}

	return m, nil
}
