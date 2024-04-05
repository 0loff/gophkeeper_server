package postgres

import (
	"context"
	"log"
	"time"

	"github.com/0loff/gophkeeper_server/internal/logger"
	"github.com/0loff/gophkeeper_server/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type DataRepository struct {
	dbpool *pgxpool.Pool
}

func NewDataRepository(db *pgxpool.Pool) *DataRepository {
	dr := &DataRepository{
		dbpool: db,
	}

	dr.CreateTextdataTable()
	dr.CreateCredsdataTable()

	return dr
}

func (r *DataRepository) CreateTextdataTable() {
	_, err := r.dbpool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS textdata (
		id serial PRIMARY KEY,
		user_id BIGINT NOT NULL,
		text TEXT NOT NULL,
		metainfo TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
		CONSTRAINT fk_users
			FOREIGN KEY(user_id)
				REFERENCES users(id)
				ON DELETE CASCADE
		);`)
	if err != nil {
		logger.Log.Error("Unable to create TEXTDATA table", zap.Error(err))
		log.Fatal(err)
	}
}

func (r *DataRepository) CreateCredsdataTable() {
	_, err := r.dbpool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS credsdata (
		id serial PRIMARY KEY,
		user_id BIGINT NOT NULL,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		metainfo TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
		CONSTRAINT fk_users
			FOREIGN KEY(user_id)
				REFERENCES users(id)
				ON DELETE CASCADE
		);`)
	if err != nil {
		logger.Log.Error("Unable to create CREDSDATA table", zap.Error(err))
		log.Fatal(err)
	}
}

func (r *DataRepository) GetTextdata(ctx context.Context, user_id int) ([]models.TextdataEntry, error) {
	var TextdataEntries []models.TextdataEntry

	rows, err := r.dbpool.Query(ctx, `SELECT id, text, metainfo FROM textdata WHERE user_id = $1`, user_id)
	if err != nil {
		logger.Log.Error("Unrecognized data from the database \n", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Textdata models.TextdataEntry
		if err := rows.Scan(&Textdata.ID, &Textdata.Text, &Textdata.Metainfo); err != nil {
			logger.Log.Error("Unable to parse the received value", zap.Error(err))
			continue
		}

		TextdataEntries = append(TextdataEntries, Textdata)
	}

	if err = rows.Err(); err != nil {
		logger.Log.Error("Unexpected error from parse data in rows next loop", zap.Error(err))
		return nil, err
	}

	return TextdataEntries, nil
}

func (r *DataRepository) CreateTextdata(ctx context.Context, user_id int, text, metainfo string) error {
	now := time.Now()

	_, err := r.dbpool.Exec(ctx, `INSERT INTO textdata(user_id, text, metainfo, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`,
		user_id, text, metainfo, now.Format(time.RFC3339), now.Format(time.RFC3339))
	if err != nil {

		logger.Log.Error("Failed to create new text", zap.Error(err))
		return err
	}

	return nil
}

func (r *DataRepository) UpdateTextdata(ctx context.Context, id int, text, metainfo string) error {
	now := time.Now()

	if _, err := r.dbpool.Exec(
		ctx,
		`UPDATE textdata
		SET text = $1, metainfo = $2, updated_at = $3
		WHERE id = $4`,
		text, metainfo, now, id,
	); err != nil {
		logger.Log.Error("Failed to update textdata", zap.Error(err))
		return err
	}

	return nil
}

func (r *DataRepository) CreateCredsdata(ctx context.Context, user_id int, username, password, metainfo string) error {
	now := time.Now()

	_, err := r.dbpool.Exec(ctx, `INSERT INTO credsdata(user_id, username, password, metainfo, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`,
		user_id, username, password, metainfo, now.Format(time.RFC3339), now.Format(time.RFC3339))
	if err != nil {

		logger.Log.Error("Failed to create new credentials", zap.Error(err))
		return err
	}

	return nil
}

func (r *DataRepository) GetCredsdata(ctx context.Context, user_id int) ([]models.CredsdataEntry, error) {
	var CredsdataEntries []models.CredsdataEntry

	rows, err := r.dbpool.Query(ctx, `SELECT id, username, password, metainfo FROM credsdata WHERE user_id = $1`, user_id)
	if err != nil {
		logger.Log.Error("Unrecognized data from the database \n", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Credsdata models.CredsdataEntry
		if err := rows.Scan(&Credsdata.ID, &Credsdata.Username, &Credsdata.Password, &Credsdata.Metainfo); err != nil {
			logger.Log.Error("Unable to parse the received value", zap.Error(err))
			continue
		}

		CredsdataEntries = append(CredsdataEntries, Credsdata)
	}

	if err = rows.Err(); err != nil {
		logger.Log.Error("Unexpected error from parse data in rows next loop", zap.Error(err))
		return nil, err
	}

	return CredsdataEntries, nil
}

func (r *DataRepository) UpdateCredsdata(ctx context.Context, id int, username, password, metainfo string) error {
	now := time.Now()

	if _, err := r.dbpool.Exec(
		ctx,
		`UPDATE credsdata
		SET username = $1, password = $2, metainfo = $3, updated_at = $4
		WHERE id = $5`,
		username, password, metainfo, now, id,
	); err != nil {
		logger.Log.Error("Failed to update textdata", zap.Error(err))
		return err
	}

	return nil
}
