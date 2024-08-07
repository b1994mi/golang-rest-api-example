package handler

import (
	"time"

	"github.com/b1994mi/golang-rest-api-example/util"
	"github.com/uptrace/bunrouter"
)

func (h *handler) VerifyHandler(r bunrouter.Request) (any, error) {
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

	m.IsUserActive = true
	now := time.Now()
	m.VerificationAt = &now

	tx := h.userRepo.StartTx()
	defer tx.Rollback()

	err = h.userRepo.Update(m, tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return m, nil
}
