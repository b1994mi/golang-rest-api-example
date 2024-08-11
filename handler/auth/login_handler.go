package auth

import (
	"fmt"
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/b1994mi/golang-rest-api-example/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (h *handler) LoginHandler(req *request) (any, error) {
	now := time.Now()

	user, err := h.userRepo.FindOneBy(map[string]interface{}{
		"phone_number": req.PhoneNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find user with phone_number %v: %v", req.PhoneNumber, err)
	}

	if !IsCorrectPass(user.Pin, req.PIN) {
		return nil, util.New401Res("001: phone number and PIN doesn’t match")
	}

	expyDur, err := time.ParseDuration("30m") // TODO: put this on .env someday
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user_id": user.ID,
		"iat":     now.Unix(),
		"exp":     now.Add(expyDur).Unix(),
	})

	tokenString, err := token.SignedString([]byte("some-secret"))
	if err != nil {
		return nil, err
	}

	refreshToken := uuid.New().String()

	tx := h.userRepo.StartTx()
	defer tx.Rollback()

	expyDurRefreshToken, err := time.ParseDuration("336h") // TODO: put this on .env someday
	if err != nil {
		return nil, err
	}

	_, err = h.userTokenRepo.Create(&model.UserToken{
		Token:      refreshToken,
		UserID:     user.ID,
		ExpDateStr: now.Add(expyDurRefreshToken).Format(time.RFC3339),
	}, tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return map[string]any{
		"access_token":  tokenString,
		"refresh_token": refreshToken,
	}, nil
}

func IsCorrectPass(hashed string, plain string) bool {
	byteHash := []byte(hashed)
	bytePlain := []byte(plain)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	return err == nil
}
