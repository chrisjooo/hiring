package usecase

import (
	"time"

	"github.com/christianchrisjo/hiring/internal/models"
	"github.com/google/uuid"
)

func (u *Usecase) CreateJob(req models.CreateJobRequest) (models.Job, error) {
	err := req.Validate()
	if err != nil {
		return models.Job{}, err
	}

	job, err := u.postgres.CreateJob(models.Job{
		ID:          uuid.New(),
		CompanyName: req.CompanyName,
		Title:       req.Title,
		Description: req.Description,
		Status:      models.JobStatusHiring,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return models.Job{}, err
	}

	return job, nil
}

func (u *Usecase) GetJobByID(id string) (models.Job, error) {
	job, err := u.postgres.GetJobByID(id)
	if err != nil {
		return models.Job{}, err
	}

	return job, nil
}

func (u *Usecase) GetAllJobs() ([]models.Job, error) {
	jobs, err := u.postgres.GetAllJobs()
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (u *Usecase) UpdateJob(req models.UpdateJobRequest) (models.Job, error) {
	existing, err := u.postgres.GetJobByID(req.ID.String())
	if err != nil {
		return models.Job{}, err
	}

	err = req.Validate(existing)
	if err != nil {
		return models.Job{}, err
	}

	job, err := u.postgres.UpdateJob(models.Job{
		ID:          req.ID,
		Description: req.Description,
		Status:      req.Status,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return models.Job{}, err
	}

	return job, nil
}

func (u *Usecase) DeleteJob(id string) error {
	err := u.postgres.DeleteJob(id)
	if err != nil {
		return err
	}
	return err
}
