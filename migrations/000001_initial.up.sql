-- +goose Up

CREATE TYPE enum_user_type AS ENUM (
    'employer',
    'employee'
);

CREATE TABLE users(
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    type enum_user_type NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ
);
CREATE INDEX ON users(email);

CREATE TYPE enum_job_status AS ENUM (
    'hiring',
    'not hiring'
);

CREATE TABLE IF NOT EXISTS jobs(
    id UUID PRIMARY KEY,
    company_name VARCHAR(256) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(16) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ
);
CREATE INDEX ON jobs(company_name);

-- CREATE TYPE enum_job_application_status AS ENUM (
--     'pending',
--     'interview',
--     'accepted',
--     'rejected'
-- );

-- CREATE TABLE IF NOT EXISTS job_applications(
--     id UUID PRIMARY KEY,
--     user_id UUID NOT NULL,
--     job_id UUID NOT NULL,
--     status enum_job_application_status NOT NULL DEFAULT 'pending',
--     created_at TIMESTAMPTZ NOT NULL,
--     updated_at TIMESTAMPTZ,
--     UNIQUE (user_id, job_id)
-- );
-- CREATE INDEX on job_application(user_id);
-- CREATE INDEX on job_application(job_id);
