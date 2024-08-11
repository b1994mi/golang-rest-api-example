package transaction

import (
	"errors"
	"time"

	"github.com/b1994mi/golang-rest-api-example/message"
	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *handler) TransferConsumer(req *message.Transfer) error {
	now := time.Now()

	tx := h.userRepo.StartTx()
	defer tx.Rollback()

	user, err := h.userRepo.FindOneForUpdateBy(map[string]any{
		"id": req.TargetUser,
	}, tx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	balanceBefore := user.Wallet
	user.Wallet = user.Wallet + req.Amount
	err = h.userRepo.Update(user, tx)
	if err != nil {
		return err
	}

	_, err = h.userTransactionRepo.Create(&model.UserTransaction{
		ID:              uuid.New().String(),
		UserID:          user.ID,
		HandlingType:    model.Transfer,
		TransactionType: model.Credit,
		Status:          model.Sucess,
		Amount:          req.Amount,
		Remarks:         req.Remarks,
		BalanceBefore:   balanceBefore,
		BalanceAfter:    user.Wallet,
		CreatedAt:       now,
	}, tx)
	if err != nil {
		return err
	}

	userTransaction, err := h.userTransactionRepo.FindOneForUpdateBy(map[string]any{
		"id": req.TrfSource,
	}, tx)
	if err != nil {
		return err
	}

	userTransaction.Status = model.Sucess
	err = h.userTransactionRepo.Update(userTransaction, tx)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
