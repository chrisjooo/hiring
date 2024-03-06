package usecase

import (
	"errors"
	"time"

	"github.com/christianchrisjo/hiring/internal/models"
	"github.com/google/uuid"
)

func (u *Usecase) CreateJobApplication(req models.CreateJobApplicationRequest) (models.JobApplication, error) {
	_, err := u.GetJobByID(req.JobID.String())
	if err != nil {
		return models.JobApplication{}, errors.New("job not found")
	}

	jobApplication, err := u.postgres.CreateJobApplication(models.JobApplication{
		ID:        uuid.New(),
		JobID:     req.JobID,
		UserID:    req.UserID,
		Status:    models.JobApplicationStatusPending,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return models.JobApplication{}, err
	}

	return jobApplication, nil
}

func (u *Usecase) GetJobApplicationByID(id string) (models.JobApplication, error) {
	jobApplication, err := u.postgres.GetJobApplicationByID(id)
	if err != nil {
		return models.JobApplication{}, err
	}

	return jobApplication, nil
}

// GetJobApplicationsByJobID returns all job applications by job id (employer view)
func (u *Usecase) GetJobApplicationsByJobID(jobID string) ([]models.JobApplication, error) {
	jobApplications, err := u.postgres.GetJobApplicationsByJobID(jobID)
	if err != nil {
		return nil, err
	}

	return jobApplications, nil
}

// GetJobApplicationByUserID returns all job applications by user id (employee view)
func (u *Usecase) GetJobApplicationByUserID(userID string) ([]models.JobApplication, error) {
	jobApplications, err := u.postgres.GetJobApplicationsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return jobApplications, nil
}

// GetJobApplicationByJobIDAndUserID returns job application by job id and user id (employee and employer view)
func (u *Usecase) GetJobApplicationByJobIDAndUserID(jobID, userID string) (models.JobApplication, error) {
	jobApplication, err := u.postgres.GetJobApplicationByJobIDAndUserID(jobID, userID)
	if err != nil {
		return models.JobApplication{}, err
	}

	return jobApplication, nil
}

func (u *Usecase) UpdateJobApplication(req models.UpdateJobApplicationRequest) (models.JobApplication, error) {
	err := req.Validate()
	if err != nil {
		return models.JobApplication{}, err
	}

	jobApplication, err := u.postgres.UpdateJobApplication(models.JobApplication{
		ID:     req.ID,
		Status: req.Status,
	})
	if err != nil {
		return models.JobApplication{}, err
	}

	return jobApplication, nil
}
