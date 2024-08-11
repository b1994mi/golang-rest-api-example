package user

import (
	"errors"

	"github.com/b1994mi/golang-rest-api-example/util"
	"gorm.io/gorm"
)

func (h *handler) FindHandler(req *request) (any, error) {
	user, err := h.userRepo.FindOneBy(map[string]interface{}{
		"id": req.ID,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user == nil {
		return nil, util.New404Res("user with id %v is not found", req.ID)
	}

	return user, nil
}
