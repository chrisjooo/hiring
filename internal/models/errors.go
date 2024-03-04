package models

import "errors"

var (
	ErrIDRequired          = errors.New("id is required")
	ErrEmailRequired       = errors.New("email is required")
	ErrPasswordRequired    = errors.New("password is required")
	ErrTypeRequired        = errors.New("type is required")
	ErrDescriptionRequired = errors.New("description is required")
	ErrCompanyNameRequired = errors.New("company name is required")
	ErrInvalidType         = errors.New("invalid user type")
)
