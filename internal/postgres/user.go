package postgres

import (
	"database/sql"

	"github.com/christianchrisjo/hiring/internal/models"
)

func (p *Postgres) CreateUser(req models.User) (models.User, error) {
	query := `INSERT INTO users (id, email, password, type, description, created_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := p.db.Exec(query, req.UserID, req.Email, req.Password, req.Type, req.Description, req.CreatedAt)
	if err != nil {
		return models.User{}, err
	}
	return req, nil
}

func (p *Postgres) GetUserByEmail(email string) (models.User, error) {
	query := `SELECT id, email, type, description, created_at, updated_at FROM users WHERE email = $1`
	user := models.User{}
	row := p.db.QueryRow(query, email)

	var updatedAt sql.NullTime

	err := row.Scan(
		&user.UserID,
		&user.Email,
		&user.Type,
		&user.Description,
		&user.CreatedAt,
		&updatedAt)
	if err != nil {
		return models.User{}, err
	}
	user.UpdatedAt = updatedAt.Time
	return user, nil
}

func (p *Postgres) UpdateUser(req models.UpdateUserRequest) (models.User, error) {
	query := `UPDATE users SET description = $1, updated_at = time.Now() WHERE id = $2 
	RETURNING id, email, type, description, created_at, updated_at`

	row := p.db.QueryRow(query, req.Description, req.UserID)

	var user models.User
	var updatedAt sql.NullTime
	err := row.Scan(
		&user.UserID,
		&user.Email,
		&user.Type,
		&user.Description,
		&user.CreatedAt,
		&updatedAt)
	if err != nil {
		return models.User{}, err
	}
	user.UpdatedAt = updatedAt.Time
	return user, nil
}
