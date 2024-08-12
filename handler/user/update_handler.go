package user

import (
	"errors"

	"github.com/b1994mi/golang-rest-api-example/util"
	"gorm.io/gorm"
)

func (h *handler) UpdateHandler(req *request) (any, error) {
	tx := h.userRepo.StartTx()
	defer tx.Rollback()

	user, err := h.userRepo.FindOneForUpdateBy(map[string]interface{}{
		"id": req.ID,
	}, tx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user == nil {
		return nil, util.New404Res("user with id %v is not found", req.ID)
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Address = req.Address

	err = h.userRepo.Update(user, tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	user.UpdatedDate = user.UpdatedAt.Format("2006-01-02 15:04:05")
	return user, nil
}
