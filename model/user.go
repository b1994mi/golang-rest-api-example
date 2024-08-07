package model

import (
	"time"

	"github.com/b1994mi/golang-rest-api-example/util"
	"gorm.io/gorm"
)

type User struct {
	ID             int        `json:"id"`
	Email          string     `json:"email"`
	Name           string     `json:"name"`
	PhoneNumber    string     `json:"phone_number"`
	Address        string     `json:"address"`
	Password       string     `json:"password"`
	IsUserActive   bool       `json:"is_user_active"`
	VerificationAt *time.Time `json:"verification_at"`
	ProfileImage   string     `json:"profile_image"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

type UserRepo interface {
	StartTx() *gorm.DB
	Create(m *User, tx *gorm.DB) (*User, error)
	Update(m *User, tx *gorm.DB) error
	Delete(m *User, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*User, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*User, error)
	Count(criteria map[string]interface{}) int64
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db,
	}
}

func (rpo *userRepo) StartTx() *gorm.DB {
	return rpo.db.Begin()
}

func (rpo *userRepo) Create(m *User, tx *gorm.DB) (*User, error) {
	err := tx.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (rpo *userRepo) Update(m *User, tx *gorm.DB) error {
	return tx.Save(&m).Error
}

func (rpo *userRepo) Delete(m *User, tx *gorm.DB) error {
	return tx.Delete(&m).Error
}

func (rpo *userRepo) FindOneBy(criteria map[string]interface{}) (*User, error) {
	var m User

	err := rpo.db.Where(criteria).Take(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (rpo *userRepo) FindBy(criteria map[string]interface{}, page, size int) ([]*User, error) {
	var data []*User
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

func (rpo *userRepo) Count(criteria map[string]interface{}) int64 {
	var result int64

	if res := rpo.db.Model(User{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return result
}
