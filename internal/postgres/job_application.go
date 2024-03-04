package postgres

import (
	"database/sql"

	"github.com/christianchrisjo/hiring/internal/models"
)

func (p *Postgres) CreateJobApplication(req models.JobApplication) (models.JobApplication, error) {
	query := `INSERT INTO job_applications (id, job_id, user_id, status, created_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := p.db.Exec(query, req.ID, req.JobID, req.UserID, req.Status, req.CreatedAt)
	if err != nil {
		return models.JobApplication{}, err
	}
	return req, nil
}

func (p *Postgres) GetJobApplicationByID(id string) (models.JobApplication, error) {
	query := `SELECT id, job_id, user_id, status, created_at, updated_at FROM job_applications WHERE id = $1`
	jobApplication := models.JobApplication{}
	row := p.db.QueryRow(query, id)

	var updatedAt sql.NullTime

	err := row.Scan(
		&jobApplication.ID,
		&jobApplication.JobID,
		&jobApplication.UserID,
		&jobApplication.Status,
		&jobApplication.CreatedAt,
		&updatedAt)
	if err != nil {
		return models.JobApplication{}, err
	}
	jobApplication.UpdatedAt = updatedAt.Time
	return jobApplication, nil
}

func (p *Postgres) GetJobApplicationsByJobID(jobID string) ([]models.JobApplication, error) {
	query := `SELECT id, job_id, user_id, status, created_at, updated_at FROM job_applications WHERE job_id = $1`
	jobApplications := []models.JobApplication{}

	rows, err := p.db.Query(query, jobID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var jobApplication models.JobApplication
		var updatedAt sql.NullTime
		if err := rows.Scan(
			&jobApplication.ID,
			&jobApplication.JobID,
			&jobApplication.UserID,
			&jobApplication.Status,
			&jobApplication.CreatedAt,
			&updatedAt); err != nil {
			return nil, err
		}
		jobApplication.UpdatedAt = updatedAt.Time
		jobApplications = append(jobApplications, jobApplication)
	}
	return jobApplications, nil
}

func (p *Postgres) GetJobApplicationsByUserID(userID string) ([]models.JobApplication, error) {
	query := `SELECT id, job_id, user_id, status, created_at, updated_at FROM job_applications WHERE user_id = $1`
	jobApplications := []models.JobApplication{}

	rows, err := p.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var jobApplication models.JobApplication
		var updatedAt sql.NullTime
		if err := rows.Scan(
			&jobApplication.ID,
			&jobApplication.JobID,
			&jobApplication.UserID,
			&jobApplication.Status,
			&jobApplication.CreatedAt,
			&updatedAt); err != nil {
			return nil, err
		}
		jobApplication.UpdatedAt = updatedAt.Time
		jobApplications = append(jobApplications, jobApplication)
	}
	return jobApplications, nil
}

func (p *Postgres) GetJobApplicationByJobIDAndUserID(jobID, userID string) (models.JobApplication, error) {
	query := `SELECT id, job_id, user_id, status, created_at, updated_at FROM job_applications WHERE job_id = $1 AND user_id = $2`
	jobApplication := models.JobApplication{}
	row := p.db.QueryRow(query, jobID, userID)

	var updatedAt sql.NullTime

	err := row.Scan(
		&jobApplication.ID,
		&jobApplication.JobID,
		&jobApplication.UserID,
		&jobApplication.Status,
		&jobApplication.CreatedAt,
		&updatedAt)
	if err != nil {
		return models.JobApplication{}, err
	}
	return jobApplication, nil
}

func (p *Postgres) UpdateJobApplication(req models.JobApplication) (models.JobApplication, error) {
	query := `UPDATE job_applications SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := p.db.Exec(query, req.Status, req.UpdatedAt, req.ID)
	if err != nil {
		return models.JobApplication{}, err
	}
	return req, nil
}
