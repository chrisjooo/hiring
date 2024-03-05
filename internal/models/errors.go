package models

import "errors"

var (
	ErrJobIDRequired                = errors.New("job id is required")
	ErrUserIDRequired               = errors.New("user id is required")
	ErrJobApplicationIDRequired     = errors.New("job application id is required")
	ErrEmailRequired                = errors.New("email is required")
	ErrPasswordRequired             = errors.New("password is required")
	ErrTypeRequired                 = errors.New("type is required")
	ErrDescriptionRequired          = errors.New("description is required")
	ErrCompanyNameRequired          = errors.New("company name is required")
	ErrInvalidType                  = errors.New("invalid user type")
	ErrJobApplicationStatusRequired = errors.New("job application status is required")
	ErrInvalidJobApplicationStatus  = errors.New("invalid job application status")
	ErrNameRequired                 = errors.New("name is required")
	ErrInvalidCreds                 = errors.New("invalid credentials")
	ErrTokenExpired                 = errors.New("token expired")
)
