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
	dr.CreateCardsdataTable()
	dr.CreateBindataTable()

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
		username BYTEA NOT NULL,
		password BYTEA NOT NULL,
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

func (r *DataRepository) CreateCardsdataTable() {
	_, err := r.dbpool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS cardsdata (
		id serial PRIMARY KEY,
		user_id BIGINT NOT NULL,
		pan BYTEA NOT NULL,
		expiry BYTEA NOT NULL,
		holder BYTEA NOT NULL,
		metainfo TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
		CONSTRAINT fk_users
			FOREIGN KEY(user_id)
				REFERENCES users(id)
				ON DELETE CASCADE
		);`)
	if err != nil {
		logger.Log.Error("Unable to create CARDSDATA table", zap.Error(err))
		log.Fatal(err)
	}
}

func (r *DataRepository) CreateBindataTable() {
	_, err := r.dbpool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS bindata (
		id serial PRIMARY KEY,
		user_id BIGINT NOT NULL,
		bin BYTEA NOT NULL,
		metainfo TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
		CONSTRAINT fk_users
			FOREIGN KEY(user_id)
				REFERENCES users(id)
				ON DELETE CASCADE
		);`)
	if err != nil {
		logger.Log.Error("Unable to create BINDATA table", zap.Error(err))
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

func (r *DataRepository) DeleteTextdata(ctx context.Context, id int) error {
	_, err := r.dbpool.Exec(ctx, `DELETE FROM textdata WHERE id = $1`, id)
	if err != nil {
		logger.Log.Error("Failed to delete text data", zap.Error(err))
		return err
	}

	return nil
}

func (r *DataRepository) CreateCredsdata(ctx context.Context, user_id int, username, password []byte, metainfo string) error {
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

func (r *DataRepository) UpdateCredsdata(ctx context.Context, id int, username, password []byte, metainfo string) error {
	now := time.Now()

	if _, err := r.dbpool.Exec(
		ctx,
		`UPDATE credsdata
		SET username = $1, password = $2, metainfo = $3, updated_at = $4
		WHERE id = $5`,
		username, password, metainfo, now, id,
	); err != nil {
		logger.Log.Error("Failed to update crdesdata table entry", zap.Error(err))
		return err
	}

	return nil
}

func (r *DataRepository) DeleteCredsdata(ctx context.Context, id int) error {
	_, err := r.dbpool.Exec(ctx, `DELETE FROM credsdata WHERE id = $1`, id)
	if err != nil {
		logger.Log.Error("Failed to delete credentials data", zap.Error(err))
		return err
	}

	return nil
}

func (r *DataRepository) CreateCardsdata(ctx context.Context, user_id int, pan, expiry, holder []byte, metainfo string) error {
	now := time.Now()

	_, err := r.dbpool.Exec(ctx, `INSERT INTO cardsdata(user_id, pan, expiry, holder, metainfo, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		user_id, pan, expiry, holder, metainfo, now.Format(time.RFC3339), now.Format(time.RFC3339))
	if err != nil {

		logger.Log.Error("Failed to create new card data", zap.Error(err))
		return err
	}

	return nil
}

func (r *DataRepository) GetCardsdata(ctx context.Context, user_id int) ([]models.CardsdataEntry, error) {
	var CardsdataEntries []models.CardsdataEntry

	rows, err := r.dbpool.Query(ctx, `SELECT id, pan, expiry, holder, metainfo FROM cardsdata WHERE user_id = $1`, user_id)
	if err != nil {
		logger.Log.Error("Unrecognized data from the database \n", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Cardsdata models.CardsdataEntry
		if err := rows.Scan(&Cardsdata.ID, &Cardsdata.Pan, &Cardsdata.Expiry, &Cardsdata.Holder, &Cardsdata.Metainfo); err != nil {
			logger.Log.Error("Unable to parse the received value", zap.Error(err))
			continue
		}

		CardsdataEntries = append(CardsdataEntries, Cardsdata)
	}

	if err = rows.Err(); err != nil {
		logger.Log.Error("Unexpected error from parse data in rows next loop", zap.Error(err))
		return nil, err
	}

	return CardsdataEntries, nil
}

func (r *DataRepository) UpdateCardsdata(ctx context.Context, id int, pan, expiry, holder []byte, metainfo string) error {
	now := time.Now()

	if _, err := r.dbpool.Exec(
		ctx,
		`UPDATE cardsdata
		SET pan = $1, expiry = $2, holder = $3, metainfo = $4, updated_at = $5
		WHERE id = $6`,
		pan, expiry, holder, metainfo, now, id,
	); err != nil {
		logger.Log.Error("Failed to update cardsdata table entry", zap.Error(err))
		return err
	}

	return nil
}

func (r *DataRepository) DeleteCardsdata(ctx context.Context, id int) error {
	_, err := r.dbpool.Exec(ctx, `DELETE FROM cardsdata WHERE id = $1`, id)
	if err != nil {
		logger.Log.Error("Failed to delete user card data", zap.Error(err))
		return err
	}

	return nil
}

func (r *DataRepository) CreateBindata(ctx context.Context, user_id int, binary []byte, metainfo string) error {
	now := time.Now()

	_, err := r.dbpool.Exec(ctx, `INSERT INTO bindata(user_id, bin, metainfo, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`,
		user_id, binary, metainfo, now.Format(time.RFC3339), now.Format(time.RFC3339))
	if err != nil {

		logger.Log.Error("Failed to create new binary data entry", zap.Error(err))
		return err
	}

	return nil
}

func (r *DataRepository) GetBindata(ctx context.Context, user_id int) ([]models.BindataEntry, error) {
	var BindataEntries []models.BindataEntry

	rows, err := r.dbpool.Query(ctx, `SELECT id, bin, metainfo FROM bindata WHERE user_id = $1`, user_id)
	if err != nil {
		logger.Log.Error("Unrecognized data from the database \n", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Bindata models.BindataEntry
		if err := rows.Scan(&Bindata.ID, &Bindata.Binary, &Bindata.Metainfo); err != nil {
			logger.Log.Error("Unable to parse the received value", zap.Error(err))
			continue
		}

		BindataEntries = append(BindataEntries, Bindata)
	}

	if err = rows.Err(); err != nil {
		logger.Log.Error("Unexpected error from parse data in rows next loop", zap.Error(err))
		return nil, err
	}

	return BindataEntries, nil
}

func (r *DataRepository) UpdateBindata(ctx context.Context, id int, binary []byte, metainfo string) error {
	now := time.Now()

	if _, err := r.dbpool.Exec(
		ctx,
		`UPDATE bindata
		SET bin = $1, metainfo = $2, updated_at = $3
		WHERE id = $4`,
		binary, metainfo, now, id,
	); err != nil {
		logger.Log.Error("Failed to update bindata table entry", zap.Error(err))
		return err
	}

	return nil
}

func (r *DataRepository) DeleteBindata(ctx context.Context, id int) error {
	_, err := r.dbpool.Exec(ctx, `DELETE FROM bindata WHERE id = $1`, id)
	if err != nil {
		logger.Log.Error("Failed to delete user binary data", zap.Error(err))
		return err
	}

	return nil
}
