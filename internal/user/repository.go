package user

import (
	"context"
	"errors"

	"github.com/0loff/gophkeeper_server/internal/models"
)

var (
	ErrUserEmail  = errors.New("email already used")
	ErrWrongCreds = errors.New("wrong email or password")
)

type UserManager interface {
	Create(ctx context.Context, u *models.UserAuth) (string, error)
	GetIDByUUID(ctx context.Context, uuid string) (int, error)
	GetUserByEmail(ctx context.Context, email string) (*models.UserLogin, error)
}
