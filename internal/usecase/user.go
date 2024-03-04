package usecase

import (
	"time"

	"github.com/christianchrisjo/hiring/internal/models"
)

func (u *Usecase) CreateUser(req models.CreateUserRequest) (models.User, error) {
	err := req.Validate()
	if err != nil {
		return models.User{}, err
	}

	user, err := u.postgres.CreateUser(models.User{
		UserID:      req.ID,
		Email:       req.Email,
		Password:    req.Password,
		Type:        req.Type,
		Description: req.Description,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
