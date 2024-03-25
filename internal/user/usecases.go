package user

import (
	"context"
)

type UserProcessor interface {
	Auth(ctx context.Context, login, password, email string) (string, error)
	Login(ctx context.Context, email, password string) (string, error)
	GetUserID(ctx context.Context, uuid string) (int, error)
}
