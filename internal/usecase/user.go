package usecase

import (
	"time"

	"github.com/christianchrisjo/hiring/internal/models"
	"github.com/google/uuid"
)

func (u *Usecase) CreateUser(req models.CreateUserRequest) (models.User, error) {
	err := req.Validate()
	if err != nil {
		return models.User{}, err
	}

	user, err := u.postgres.CreateUser(models.User{
		UserID:      uuid.New(),
		Email:       req.Email,
		Password:    req.Password,
		Type:        req.Type,
		Description: req.Description,
		Name:        req.Name,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *Usecase) GetUserByEmail(email string) (models.User, error) {
	if email == "" {
		return models.User{}, models.ErrEmailRequired
	}
	user, err := u.postgres.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *Usecase) UpdateUser(req models.UpdateUserRequest) (models.User, error) {
	existing, err := u.postgres.GetUserByID(req.UserID.String())
	if err != nil {
		return models.User{}, err
	}

	err = req.Validate(existing)
	if err != nil {
		return models.User{}, err
	}

	// update data
	existing.Description = req.Description
	existing.Name = req.Name
	existing.UpdatedAt = time.Now()

	user, err := u.postgres.UpdateUser(existing)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
