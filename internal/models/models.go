package models

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type UserType string

const (
	Employer UserType = "employer"
	Employee UserType = "employee"
)

type User struct {
	UserID      uuid.UUID
	Email       string
	Password    string
	Type        UserType
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateUserRequest struct {
	ID          uuid.UUID
	Email       string
	Password    string
	Type        UserType
	Description string
}

func (req *CreateUserRequest) Validate() error {
	if req.ID == uuid.Nil {
		return ErrIDRequired
	}
	if req.Email == "" {
		return ErrEmailRequired
	}
	if req.Password == "" {
		return ErrPasswordRequired
	}
	if req.Type == "" {
		return ErrTypeRequired
	}
	if req.Description == "" {
		return ErrDescriptionRequired
	}

	hasher := sha256.New().Sum([]byte(req.Password))
	req.Password = hex.EncodeToString(hasher)
	return nil
}
