package data

import (
	"context"

	"github.com/0loff/gophkeeper_server/internal/models"
)

type DataProcessor interface {
	StoreTextdata(ctx context.Context, uuid int, title, article string) error
	ReceiveTextdata(ctx context.Context, uid int) []models.TextdataEntry
	UpdTextdata(ctx context.Context, id int, title, text string) error
	DelTextdata(ctx context.Context, id int) error

	StoreCredsdata(ctx context.Context, uid int, username, password, metainfo string) error
	ReceiveCredsdata(ctx context.Context, uid int) []models.CredsdataEntry
	UpdCredsdata(ctx context.Context, uid int, username, password, metainfo string) error
	DelCredsdata(ctx context.Context, id int) error

	StoreCardsdata(ctx context.Context, uid int, pan, expiry, holder, metainfo string) error
	ReceiveCardsdata(ctx context.Context, uid int) []models.CardsdataEntry
	UpdCardsdata(ctx context.Context, uid int, pan, expiry, holder, metainfo string) error
	DelCardsdata(ctx context.Context, id int) error

	StoreBindata(ctx context.Context, uid int, binary []byte, metainfo string) error
	ReceiveBindata(ctx context.Context, uid int) []models.BindataEntry
	UpdBindata(ctx context.Context, uid int, binary []byte, metainfo string) error
	DelBindata(ctx context.Context, id int) error
}
