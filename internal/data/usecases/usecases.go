package usecases

import (
	"context"

	"github.com/0loff/gophkeeper_server/internal/data"
	"github.com/0loff/gophkeeper_server/internal/logger"
	"github.com/0loff/gophkeeper_server/internal/models"
	"go.uber.org/zap"
)

type DataUseCases struct {
	ac data.DataManager
}

func NewDataUseCases(ac data.DataManager) *DataUseCases {

	return &DataUseCases{
		ac: ac,
	}
}

func (d *DataUseCases) ReceiveTextdata(ctx context.Context, uid int) []models.TextdataEntry {
	textData, err := d.ac.GetTextdata(ctx, uid)
	if err != nil {
		logger.Log.Error("Cannot receive user text data", zap.Error(err))
		return nil
	}
	return textData
}

func (d *DataUseCases) StoreTextdata(ctx context.Context, uid int, title, article string) error {
	if err := d.ac.CreateTextdata(ctx, uid, title, article); err != nil {
		logger.Log.Error("Cannot processed article creation", zap.Error(err))
		return err
	}

	return nil
}

func (d *DataUseCases) UpdTextdata(ctx context.Context, id int, title, text string) error {
	if err := d.ac.UpdateTextdata(ctx, id, title, text); err != nil {
		logger.Log.Error("Cannot processed textdata updating", zap.Error(err))
		return err
	}

	return nil
}

func (d *DataUseCases) StoreCredsdata(ctx context.Context, uid int, username, password, metainfo string) error {
	if err := d.ac.CreateCredsdata(ctx, uid, username, password, metainfo); err != nil {
		logger.Log.Error("Cannot processed credentials creation", zap.Error(err))
		return err
	}

	return nil
}
