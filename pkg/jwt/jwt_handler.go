package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/0loff/gophkeeper_server/internal/logger"
)

var ErrInvalidAccessToken = errors.New("invalid access token")

// Claims структура для хранения утверждений, входящих в состав JWT токена
type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
	Key    []byte
}

const tokenExp = time.Hour * 3
const secretKey = "secretkey"

// Конструктор, создающий JWT token
func BuildToken(uid string, key []byte) (string, error) {
	user_id, err := uuid.Parse(uid)
	if err != nil {
		logger.Log.Error("Cannot convert string to uuid value", zap.Error(err))
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		Claims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExp)),
			},
			UserID: user_id,
			Key:    key,
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Метод получения ID пользователя из JWT токена, полученного из авторизационных cookies
func ParseToken(authtoken string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(authtoken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		logger.Log.Error("The value of the authentication token could not be parsed", zap.Error(err))
		return "", err
	}

	if !token.Valid {
		logger.Log.Error("Invalid auth token param", zap.Error(err))
		return "", err
	}

	return claims.UserID.String(), nil
}
