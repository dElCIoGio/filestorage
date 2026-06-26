package db

import (
	"context"
	"database/sql"
	"time"
)

func Connect(ctx context.Context, dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
