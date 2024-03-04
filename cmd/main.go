package main

import (
	"context"

	"github.com/christianchrisjo/hiring/internal/models"
	"github.com/christianchrisjo/hiring/internal/postgres"
	"github.com/christianchrisjo/hiring/internal/usecase"
	"github.com/google/uuid"
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
	_, err = usecase.CreateUser(models.CreateUserRequest{
		ID:          uuid.New(),
		Email:       "christian@gmail.com",
		Password:    "password",
		Type:        models.Employee,
		Description: "user description",
	})
	if err != nil {
		panic(err)
	}

	defer db.Close()

	return
}
