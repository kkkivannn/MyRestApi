package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	Salt       = "dslkj932q90jdqos0219jd3fjreasokcmnurn4875678f"
	SigningKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	TokenTTL   = 1 * time.Minute
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
type Token struct {
	Token string `json:"token"`
}

type ManagerInterface interface {
	AccessTokenGenerate(userId int) (string, error)
}

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) NewTokenGenerate(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})
	return token.SignedString([]byte(SigningKey))
}

func (m *Manager) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(SigningKey), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("Token not valid")
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("Token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
