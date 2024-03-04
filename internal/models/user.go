package models

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type UserType string

const (
	UserTypeEmployer UserType = "employer"
	UserTypeEmployee UserType = "employee"
)

type User struct {
	UserID      uuid.UUID `json:"user_id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Type        UserType  `json:"user_type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	ID          uuid.UUID `json:"user_id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Type        UserType  `json:"type"`
	Description string    `json:"description"`
}

func (req *CreateUserRequest) Validate() error {
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
	switch req.Type {
	case UserTypeEmployer, UserTypeEmployee:
		return nil
	default:
		return ErrInvalidType
	}
}

type UpdateUserRequest struct {
	UserID      uuid.UUID `json:"user_id"`
	Description string    `json:"description"`
}

func (req *UpdateUserRequest) Validate() error {
	if req.UserID == uuid.Nil {
		return ErrIDRequired
	}
	if req.Description == "" {
		return ErrDescriptionRequired
	}
	return nil
}
