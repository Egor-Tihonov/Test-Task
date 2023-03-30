package db

import (
	"context"
	"time"

	"github.com/Egor-Tihonov/Test-Task/pkg/models"
	"github.com/jackc/pgx/v4"
)

func (d *DB) Get(ctx context.Context, username string) (*time.Time, error) {
	var date time.Time
	sql := "select first_req from weatherbot where username = $1"
	err := d.Pool.QueryRow(ctx, sql, username).Scan(&date)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, models.ErrorUserDontExist
		} else {
			return nil, err
		}
	}

	return &date, nil
}

func (d *DB) CreateUser(ctx context.Context, username string) error {
	sql := "insert into weatherbot(username,first_req) values ($1,$2)"
	_, err := d.Pool.Exec(ctx, sql, username, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) UpdateTime(ctx context.Context, date time.Time) error {
	sql := "update weatherbot set first_req = $1"
	_, err := d.Pool.Exec(ctx, sql, date)
	if err != nil {
		return err
	}
	return nil
}
