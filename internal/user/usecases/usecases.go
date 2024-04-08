package usecases

import (
	"context"
	"crypto/aes"
	"crypto/rand"

	"github.com/0loff/gophkeeper_server/internal/logger"
	"github.com/0loff/gophkeeper_server/internal/models"
	"github.com/0loff/gophkeeper_server/internal/user"
	"github.com/0loff/gophkeeper_server/pkg/encryptor"
	"github.com/0loff/gophkeeper_server/pkg/jwt"
	"go.uber.org/zap"
)

type UserUseCases struct {
	um user.UserManager
}

func NewUserUseCases(um user.UserManager) *UserUseCases {

	return &UserUseCases{
		um: um,
	}
}

func (u *UserUseCases) Auth(ctx context.Context, username, password, email string) (string, error) {
	hash, err := encryptor.Encrypt(password)
	if err != nil {
		logger.Log.Error("Failed to create hash from password", zap.Error(err))
		return "", err
	}

	key, err := generateRandom(2 * aes.BlockSize) // ключ шифрования
	if err != nil {
		logger.Log.Error("Failed to create uniq user key from password", zap.Error(err))
		return "", err
	}

	newUser := models.UserAuth{
		Username: username,
		Password: hash,
		Email:    email,
		Key:      key,
	}

	uuid, err := u.um.Create(ctx, &newUser)
	if err != nil {
		logger.Log.Error("Error creating a new user", zap.Error(err))
		return "", err
	}

	return jwt.BuildToken(uuid, newUser.Key)
}

func (u *UserUseCases) Login(ctx context.Context, email, password string) (string, error) {
	userEntry, err := u.um.GetUserByEmail(ctx, email)
	if err != nil {
		logger.Log.Error("Cannot find user by email.", zap.Error(err))
		return "", err
	}

	err = encryptor.Compare(userEntry.Password, password)
	if err != nil {
		logger.Log.Error("Wrong password", zap.Error(err))
		return "", user.ErrWrongCreds
	}

	return jwt.BuildToken(userEntry.UUID, userEntry.Key)
}

func (u *UserUseCases) GetUserID(ctx context.Context, uuid string) (int, error) {
	id, err := u.um.GetIDByUUID(ctx, uuid)
	if err != nil {
		logger.Log.Error("Can't find user ID", zap.Error(err))
		return 0, err
	}

	return id, nil

}

func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
