package postgres

import (
	"context"
	"database/sql"

	"github.com/christianchrisjo/hiring/internal/models"
)

func (p *Postgres) CreateJob(req models.Job) (models.Job, error) {
	query := `INSERT INTO jobs (id, company_name, title, description, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := p.db.Exec(query, req.ID, req.CompanyName, req.Title, req.Description, req.Status, req.CreatedAt)
	if err != nil {
		return models.Job{}, err
	}
	return req, nil
}

func (p *Postgres) GetJobByID(id string) (models.Job, error) {
	query := `SELECT id, company_name, title, description, status, created_at, updated_at FROM jobs WHERE id = $1`
	job := models.Job{}
	row := p.db.QueryRow(query, id)

	var updatedAt sql.NullTime
	err := row.Scan(
		&job.ID,
		&job.CompanyName,
		&job.Title,
		&job.Description,
		&job.Status,
		&job.CreatedAt,
		&updatedAt)
	if err != nil {
		return models.Job{}, err
	}
	job.UpdatedAt = updatedAt.Time
	return job, nil
}

func (p *Postgres) GetAllJobs() ([]models.Job, error) {
	query := `SELECT id, company_name, title, description, status, created_at, updated_at FROM jobs`
	jobs := []models.Job{}

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var job models.Job
		var updatedAt sql.NullTime
		if err := rows.Scan(
			&job.ID,
			&job.CompanyName,
			&job.Title,
			&job.Description,
			&job.Status,
			&job.CreatedAt,
			&updatedAt); err != nil {
			return nil, err
		}
		job.UpdatedAt = updatedAt.Time
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (p *Postgres) UpdateJob(req models.Job) (models.Job, error) {
	query := `UPDATE jobs SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5`
	_, err := p.db.Exec(query, req.Title, req.Description, req.Status, req.UpdatedAt, req.ID)
	if err != nil {
		return models.Job{}, err
	}
	return req, nil
}

func (p *Postgres) DeleteJob(id string) error {
	query := `DELETE FROM jobs WHERE id = $1`

	tx, err := p.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}
	err = deleteJobApplicationByJobID(tx, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}
