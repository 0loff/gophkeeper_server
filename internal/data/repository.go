package data

import (
	"context"

	"github.com/0loff/gophkeeper_server/internal/models"
)

type DataManager interface {
	GetTextdata(ctx context.Context, user_id int) ([]models.TextdataEntry, error)
	CreateTextdata(ctx context.Context, user_id int, title, article string) error
	UpdateTextdata(ctx context.Context, id int, title, text string) error
	DeleteTextdata(ctx context.Context, id int) error

	CreateCredsdata(ctx context.Context, user_id int, username, password, metainfo string) error
	GetCredsdata(ctx context.Context, user_id int) ([]models.CredsdataEntry, error)
	UpdateCredsdata(ctx context.Context, user_id int, username, password, metainfo string) error
	DeleteCredsdata(ctx context.Context, id int) error

	CreateCardsdata(ctx context.Context, user_id int, pan, expiry, holder, metainfo string) error
	GetCardsdata(ctx context.Context, user_id int) ([]models.CardsdataEntry, error)
	UpdateCardsdata(ctx context.Context, user_id int, pan, expiry, holder, metainfo string) error
	DeleteCardsdata(ctx context.Context, id int) error

	CreateBindata(ctx context.Context, user_id int, binary []byte, metainfo string) error
	GetBindata(ctx context.Context, user_id int) ([]models.BindataEntry, error)
	UpdateBindata(ctx context.Context, user_id int, binary []byte, metainfo string) error
	DeleteBindata(ctx context.Context, id int) error
}
