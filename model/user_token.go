package model

import (
	"time"

	"github.com/b1994mi/golang-rest-api-example/util"
	"gorm.io/gorm"
)

type UserToken struct {
	ID     int    `json:"id"`
	UserID string `json:"user_id"`
	Token  string `json:"token"`
	// use string of format YYYY-MM-DDTHH:mm:ssZ because string comparison
	// is much simpler as I've experienced while developing Ecomsyst for ACSET
	ExpDateStr string    `json:"exp_date_str"`
	CreatedAt  time.Time `json:"created_at"`
}

func (m *UserToken) TableName() string {
	return "user_tokens"
}

type UserTokenRepo interface {
	StartTx() *gorm.DB
	Create(m *UserToken, tx *gorm.DB) (*UserToken, error)
	Update(m *UserToken, tx *gorm.DB) error
	Delete(m *UserToken, tx *gorm.DB) error
	FindOneBy(criteria map[string]any) (*UserToken, error)
	FindOneWithUserBy(criteria map[string]any) (*UserTokenWithUser, error)
	FindBy(criteria map[string]any, page, size int) ([]*UserToken, error)
	Count(criteria map[string]any) int64
}

type userTokenRepo struct {
	db *gorm.DB
}

func NewUserTokenRepo(db *gorm.DB) UserTokenRepo {
	return &userTokenRepo{
		db,
	}
}

func (rpo *userTokenRepo) StartTx() *gorm.DB {
	return rpo.db.Begin()
}

func (rpo *userTokenRepo) Create(m *UserToken, tx *gorm.DB) (*UserToken, error) {
	err := tx.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (rpo *userTokenRepo) Update(m *UserToken, tx *gorm.DB) error {
	return tx.Save(&m).Error
}

func (rpo *userTokenRepo) Delete(m *UserToken, tx *gorm.DB) error {
	return tx.Delete(&m).Error
}

func (rpo *userTokenRepo) FindOneBy(criteria map[string]any) (*UserToken, error) {
	var m UserToken

	err := rpo.db.Where(criteria).Take(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

type UserTokenWithUser struct {
	*UserToken
	User *User
}

func (rpo *userTokenRepo) FindOneWithUserBy(criteria map[string]any) (*UserTokenWithUser, error) {
	var m UserTokenWithUser

	err := rpo.db.Joins("User").Where(criteria).Take(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (rpo *userTokenRepo) FindBy(criteria map[string]any, page, size int) ([]*UserToken, error) {
	var data []*UserToken
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

func (rpo *userTokenRepo) Count(criteria map[string]any) int64 {
	var result int64

	if res := rpo.db.Model(User{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return result
}
