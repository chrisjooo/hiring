package models

import (
	"time"

	"github.com/google/uuid"
)

type JobStatus string

const (
	JobStatusHiring    JobStatus = "hiring"
	JobStatusNotHiring JobStatus = "not hiring"
)

type Job struct {
	ID          uuid.UUID `json:"job_id"`
	CompanyName string    `json:"company_name"`
	Description string    `json:"description"`
	Status      JobStatus `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateJobRequest struct {
	ID          uuid.UUID `json:"job_id"`
	CompanyName string    `json:"company_name"`
	Description string    `json:"description"`
}

func (req *CreateJobRequest) Validate() error {
	if req.CompanyName == "" {
		return ErrCompanyNameRequired
	}
	if req.Description == "" {
		return ErrDescriptionRequired
	}
	return nil
}

type UpdateJobRequest struct {
	ID          uuid.UUID `json:"job_id"`
	CompanyName string    `json:"company_name"`
	Description string    `json:"description"`
	Status      JobStatus `json:"status"`
}

func (req *UpdateJobRequest) Validate(existing Job) error {
	if req.ID == uuid.Nil {
		return ErrIDRequired
	}
	if req.CompanyName == "" {
		req.CompanyName = existing.CompanyName
	}
	if req.Description == "" {
		req.Description = existing.Description
	}
	if req.Status == "" {
		req.Status = existing.Status
	}
	return nil
}
