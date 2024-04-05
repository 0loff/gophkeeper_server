package data

import (
	"context"

	"github.com/0loff/gophkeeper_server/internal/models"
)

type DataManager interface {
	GetTextdata(ctx context.Context, user_id int) ([]models.TextdataEntry, error)
	CreateTextdata(ctx context.Context, user_id int, title, article string) error
	UpdateTextdata(ctx context.Context, id int, title, text string) error

	CreateCredsdata(ctx context.Context, user_id int, username, password, metainfo string) error
}
