package auth

import (
	"errors"
	"time"

	"github.com/b1994mi/golang-rest-api-example/util"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func (h *handler) RefreshTokenHandler(req *request) (any, error) {
	now := time.Now()

	rt, err := h.userTokenRepo.FindOneWithUserBy(map[string]any{
		"token": req.RefreshToken,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if rt == nil ||
		rt.ExpDateStr < now.Format(time.RFC3339) ||
		rt.User.IsDeleted {
		return nil, util.New401Res("refresh token is not valid")
	}

	expyDur, err := time.ParseDuration("30m") // TODO: put this on .env someday
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user_id": rt.User.ID,
		"iat":     now.Unix(),
		"exp":     now.Add(expyDur).Unix(),
	})

	tokenString, err := token.SignedString([]byte("some-secret"))
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"access_token": tokenString,
	}, nil
}
