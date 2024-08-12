package transaction

import (
	"errors"
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *handler) PaymentHandler(req *request) (any, error) {
	now := time.Now()

	tx := h.userRepo.StartTx()
	defer tx.Rollback()

	user, err := h.userRepo.FindOneForUpdateBy(map[string]any{
		"id": req.UserID,
	}, tx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	balanceBefore := user.Wallet
	user.Wallet = user.Wallet - req.Amount
	err = h.userRepo.Update(user, tx)
	if err != nil {
		return nil, err
	}

	m, err := h.userTransactionRepo.Create(&model.UserTransaction{
		ID:              uuid.New().String(),
		UserID:          user.ID,
		HandlingType:    model.Payment,
		TransactionType: model.Debit,
		Status:          model.Sucess,
		Amount:          req.Amount,
		Remarks:         req.Remarks,
		BalanceBefore:   balanceBefore,
		BalanceAfter:    user.Wallet,
		CreatedAt:       now,
	}, tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	m.PaymentID = m.ID
	m.CreatedDate = m.CreatedAt.Format("2006-01-02 15:04:05")
	return m, nil
}
