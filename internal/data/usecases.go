package data

import (
	"context"

	"github.com/0loff/gophkeeper_server/internal/models"
)

type DataProcessor interface {
	StoreTextdata(ctx context.Context, uuid int, title, article string) error
	ReceiveTextdata(ctx context.Context, uid int) []models.TextdataEntry
	UpdTextdata(ctx context.Context, id int, title, text string) error
}
