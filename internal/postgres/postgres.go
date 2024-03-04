package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type Postgres struct {
	db *sql.DB
}

// New postgres init
func InitDB(ctx context.Context) (db *sql.DB, err error) {
	if err := goose.SetDialect("postgres"); err != nil {
		err = errors.New("failed to set goose dialect: " + err.Error())
	}

	connectionString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		"localhost", 5432, "hiring", "postgres", "postgres")
	db, err = sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	if err := goose.RunContext(context.Background(), "up", db, "./../migrations"); err != nil {
		return nil, errors.New("goose run: " + err.Error())
	}

	return db, nil
}

func New(db *sql.DB) (*Postgres, error) {
	if db == nil {
		return nil, errors.New("DB is nil")
	}
	return &Postgres{
		db: db,
	}, nil
}
