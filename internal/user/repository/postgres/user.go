package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/0loff/gophkeeper_server/internal/logger"
	"github.com/0loff/gophkeeper_server/internal/models"
	userpkg "github.com/0loff/gophkeeper_server/internal/user"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserRepository struct {
	dbpool *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	ur := &UserRepository{
		dbpool: db,
	}

	ur.CreateTable()
	return ur
}

func (r *UserRepository) CreateTable() {
	_, err := r.dbpool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		uuid text NOT NULL,
		username TEXT NOT NULL,
		hash TEXT NOT NULL,
		email TEXT NOT NULL,
		key BYTEA NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL
		);`)
	if err != nil {
		logger.Log.Error("Unable to create USERS table", zap.Error(err))
	}

	_, err = r.dbpool.Exec(context.Background(), "CREATE UNIQUE INDEX IF NOT EXISTS email ON users (email)")
	if err != nil {
		logger.Log.Error("Unable to create unique index for email field")
	}

	_, err = r.dbpool.Exec(context.Background(), "CREATE UNIQUE INDEX IF NOT EXISTS uuid ON users (uuid)")
	if err != nil {
		logger.Log.Error("Unable to create unique index for uuid field")
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.UserAuth) (string, error) {
	now := time.Now()
	uid := uuid.New().String()

	_, err := r.dbpool.Exec(ctx, `INSERT INTO users(uuid, username, hash, email, key, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		uid, user.Username, user.Password, user.Email, user.Key, now.Format(time.RFC3339), now.Format(time.RFC3339))
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			logger.Log.Error("Email already used", zap.Error(err))
			return "", userpkg.ErrUserEmail
		}
		logger.Log.Error("Failed to create new user", zap.Error(err))
		return "", err
	}

	return uid, err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.UserLogin, error) {
	var User models.UserLogin

	row := r.dbpool.QueryRow(ctx, `SELECT uuid, hash, key FROM users WHERE email = $1`, email)
	if err := row.Scan(&User.UUID, &User.Password, &User.Key); err != nil {
		return nil, userpkg.ErrWrongCreds
	}

	return &User, nil
}

func (r *UserRepository) GetIDByUUID(ctx context.Context, uid string) (int, error) {
	row := r.dbpool.QueryRow(ctx, `SELECT id FROM users WHERE uuid = $1`, uid)

	var id int
	if err := row.Scan(&id); err != nil {
		logger.Log.Error("Unable to parse the received id by uuid from DB", zap.Error(err))
		return 0, err
	}

	return id, nil
}
