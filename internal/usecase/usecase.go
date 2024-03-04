package usecase

import "github.com/christianchrisjo/hiring/internal/postgres"

type Usecase struct {
	postgres *postgres.Postgres
}

func New(db *postgres.Postgres) *Usecase {
	return &Usecase{
		postgres: db,
	}
}
