package usecases

import (
	"context"

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

	newUser := models.UserAuth{
		Username: username,
		Password: hash,
		Email:    email,
	}

	uuid, err := u.um.Create(ctx, &newUser)
	if err != nil {
		logger.Log.Error("Error creating a new user", zap.Error(err))
		return "", err
	}

	return jwt.BuildToken(uuid)
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

	return jwt.BuildToken(userEntry.UUID)
}

func (u *UserUseCases) GetUserID(ctx context.Context, uuid string) (int, error) {
	id, err := u.um.GetIDByUUID(ctx, uuid)
	if err != nil {
		logger.Log.Error("Can't find user ID", zap.Error(err))
		return 0, err
	}

	return id, nil

}

// func (u *UserUseCases) Login(ctx context.Context, username, password string) (string, error) {
// 	hash, err := encryptor.Encrypt(password)
// 	if err != nil {
// 		logger.Log.Error("Failed to create hash from password", zap.Error(err))
// 	}

// 	newUser := &models.UserLogin{
// 		Username: username,
// 		Password: hash,
// 	}

// 	uid, err := u.uc.Create(ctx, newUser)
// 	if err != nil {
// 		logger.Log.Error("Error creating a new user", zap.Error(err))
// 		return "", err
// 	}
// 	//TODO change "secretKey" => dynamic received value
// 	return jwt.BuildToken(uid, "secretKey")
// }
