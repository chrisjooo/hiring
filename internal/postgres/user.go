package postgres

import (
	"fmt"

	"github.com/christianchrisjo/hiring/internal/models"
)

func (p *Postgres) CreateUser(req models.User) (models.User, error) {
	query := `INSERT INTO users (id, email, password, type, description, created_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := p.db.Exec(query, req.UserID, req.Email, req.Password, req.Type, req.Description, req.CreatedAt)
	if err != nil {
		fmt.Println("sampe sini")
		return models.User{}, err
	}
	return req, nil
}
