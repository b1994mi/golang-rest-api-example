package transaction

import (
	"github.com/b1994mi/golang-rest-api-example/model"
)

func (h *handler) GetHandler(req *request) (any, error) {
	// TODO: implement pagination
	transactions, err := h.userTransactionRepo.FindBy(map[string]any{
		"user_id": req.UserID,
	}, 0, 0, "created_at desc")
	if err != nil {
		return nil, err
	}

	for _, t := range transactions {
		switch t.HandlingType {
		case model.Payment:
			t.PaymentID = t.ID
		case model.TopUp:
			t.TopUpID = t.ID
		case model.Transfer:
			t.TransferID = t.ID
		}

		t.CreatedDate = t.CreatedAt.Format("2006-01-02 15:04:05")
	}

	return transactions, nil
}
