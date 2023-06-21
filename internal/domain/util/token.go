package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/a13hander/auth-service-api/internal/domain/model"
)

func GenerateToken(user *model.User, secretKey []byte, duration time.Duration) (string, error) {
	clms := model.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		Username: user.Username,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clms)

	return token.SignedString(secretKey)
}

func VerifyToken(token string, secretKey []byte) (*model.Claims, error) {
	t, err := jwt.ParseWithClaims(token, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected token singing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := t.Claims.(*model.Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
