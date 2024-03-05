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
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	Type        UserType `json:"type"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
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
	if req.Name == "" {
		return ErrNameRequired
	}

	req.Password = HashPassword(req.Password)
	switch req.Type {
	case UserTypeEmployer, UserTypeEmployee:
		return nil
	default:
		return ErrInvalidType
	}
}

func HashPassword(password string) string {
	hasher := sha256.New().Sum([]byte(password))
	return hex.EncodeToString(hasher)
}

type UpdateUserRequest struct {
	UserID      uuid.UUID `json:"user_id"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
}

func (req *UpdateUserRequest) Validate(existing User) error {
	if req.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if req.Description == "" {
		req.Description = existing.Description
	}
	if req.Name == "" {
		req.Name = existing.Name
	}

	return nil
}

type ClaimToken struct {
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	ExpiredAt int64     `json:"exp"`
}

func (c *ClaimToken) Valid() error {
	if time.Now().After(time.Unix(c.ExpiredAt, 0)) {
		return ErrTokenExpired
	}
	return nil
}

type SignInByEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *SignInByEmailRequest) Validate() error {
	if req.Email == "" {
		return ErrEmailRequired
	}
	if req.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}
