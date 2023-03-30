package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(connStr string) (*DB, error) {
	conn, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	if _, err := conn.Exec(context.Background(), "Create table IF NOT EXISTS weatherbot(username text primary key not null, first_req date )"); err != nil {
		return nil, err
	}

	return &DB{
		Pool: conn,
	}, nil
}
