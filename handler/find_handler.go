package handler

import (
	"github.com/b1994mi/golang-rest-api-example/util"
	"github.com/uptrace/bunrouter"
)

func (h *handler) FindHandler(r bunrouter.Request) (any, error) {
	var reqBody reqBody
	err := util.ShouldBindJSON(&reqBody, r)
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
