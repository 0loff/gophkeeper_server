package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(DSN string) (*pgxpool.Pool, error) {
	p, err := pgxpool.New(context.Background(), DSN)
	if err != nil {
		log.Fatal("Error occured while established connection to database", err)
	}

	connect, err := p.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from the db pool")
	}
	defer connect.Release()

	err = connect.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return p, err
}
