package models

import (
	"time"

	"github.com/google/uuid"
)

type JobApplicationStatus string

const (
	JobApplicationStatusPending   JobApplicationStatus = "pending"
	JobApplicationStatusInterview JobApplicationStatus = "interview"
	JobApplicationStatusAccepted  JobApplicationStatus = "accepted"
	JobApplicationStatusRejected  JobApplicationStatus = "rejected"
)

type JobApplication struct {
	ID        uuid.UUID            `json:"job_application_id"`
	JobID     uuid.UUID            `json:"job_id"`
	UserID    uuid.UUID            `json:"user_id"`
	Status    JobApplicationStatus `json:"status"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

type CreateJobApplicationRequest struct {
	JobID  uuid.UUID `json:"job_id"`
	UserID uuid.UUID `json:"user_id"`
}

func (req *CreateJobApplicationRequest) Validate() error {
	if req.JobID == uuid.Nil {
		return ErrJobIDRequired
	}
	if req.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	return nil
}

type UpdateJobApplicationRequest struct {
	ID     uuid.UUID            `json:"job_application_id"`
	Status JobApplicationStatus `json:"status"`
}

func (req *UpdateJobApplicationRequest) Validate() error {
	if req.ID == uuid.Nil {
		return ErrJobApplicationIDRequired
	}
	if req.Status == "" {
		return ErrJobApplicationStatusRequired
	}
	switch req.Status {
	case JobApplicationStatusPending, JobApplicationStatusInterview, JobApplicationStatusAccepted, JobApplicationStatusRejected:
		return nil
	default:
		return ErrInvalidJobApplicationStatus
	}
}
