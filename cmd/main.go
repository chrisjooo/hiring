package main

import (
	"context"

	http "github.com/christianchrisjo/hiring/cmd/http"
	"github.com/christianchrisjo/hiring/internal/postgres"
	"github.com/christianchrisjo/hiring/internal/usecase"
)

func main() {
	ctx := context.Background()

	db, err := postgres.InitDB(ctx)
	if err != nil {
		panic(err)
	}
	repo, err := postgres.New(db)
	if err != nil {
		panic(err)
	}

	usecase := usecase.New(repo)
	defer db.Close()

	handler := http.NewHandlers(usecase)
	http.HandleRequests(handler)
}
