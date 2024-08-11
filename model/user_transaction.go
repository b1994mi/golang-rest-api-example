package model

import (
	"time"

	"github.com/b1994mi/golang-rest-api-example/util"
	"gorm.io/gorm"
)

type UserTransaction struct {
	ID              string            `json:"-"`
	UserID          string            `json:"user_id"`
	HandlingType    HandlingType      `json:"handling_type"`
	TransactionType TransactionType   `json:"transaction_type"`
	Status          TransactionStatus `json:"status"`
	Amount          float64           `json:"amount"`
	Remarks         string            `json:"remarks"`
	BalanceBefore   float64           `json:"balance_before"`
	BalanceAfter    float64           `json:"balance_after"`
	CreatedAt       time.Time         `json:"created_at"`

	TopUpID    string `json:"top_up_id" gorm:"-"`
	PaymentID  string `json:"payment_id" gorm:"-"`
	TransferID string `json:"transfer_id" gorm:"-"`
}

func (m *UserTransaction) TableName() string {
	return "user_transactions"
}

type HandlingType string
type TransactionType string
type TransactionStatus string

const (
	TopUp    HandlingType = "TopUp"
	Payment  HandlingType = "Payment"
	Transfer HandlingType = "Transfer"

	Debit  TransactionType = "DEBIT"
	Credit TransactionType = "CREDIT"

	Sucess     TransactionStatus = "SUCCESS"
	Processing TransactionStatus = "PROCESSING"
	Failed     TransactionStatus = "FAILED"
)

type UserTransactionRepo interface {
	StartTx() *gorm.DB
	Create(m *UserTransaction, tx *gorm.DB) (*UserTransaction, error)
	Update(m *UserTransaction, tx *gorm.DB) error
	Delete(m *UserTransaction, tx *gorm.DB) error
	FindOneBy(criteria map[string]any) (*UserTransaction, error)
	FindBy(criteria map[string]any, page, size int) ([]*UserTransaction, error)
	Count(criteria map[string]any) int64
}

type userTransactionRepo struct {
	db *gorm.DB
}

func NewUserTransactionRepo(db *gorm.DB) UserTransactionRepo {
	return &userTransactionRepo{
		db,
	}
}

func (rpo *userTransactionRepo) StartTx() *gorm.DB {
	return rpo.db.Begin()
}

func (rpo *userTransactionRepo) Create(m *UserTransaction, tx *gorm.DB) (*UserTransaction, error) {
	err := tx.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (rpo *userTransactionRepo) Update(m *UserTransaction, tx *gorm.DB) error {
	return tx.Save(&m).Error
}

func (rpo *userTransactionRepo) Delete(m *UserTransaction, tx *gorm.DB) error {
	return tx.Delete(&m).Error
}

func (rpo *userTransactionRepo) FindOneBy(criteria map[string]any) (*UserTransaction, error) {
	var m *UserTransaction

	err := rpo.db.Where(criteria).Take(m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (rpo *userTransactionRepo) FindBy(criteria map[string]any, page, size int) ([]*UserTransaction, error) {
	var data []*UserTransaction
	if page == 0 || size == 0 {
		page, size = -1, -1
	}

	limit, offset := util.GetLimitOffset(page, size)
	err := rpo.db.
		Where(criteria).
		Offset(offset).Limit(limit).
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (rpo *userTransactionRepo) Count(criteria map[string]any) int64 {
	var result int64

	if res := rpo.db.Model(User{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return result
}
