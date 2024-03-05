package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/christianchrisjo/hiring/internal/models"
	"github.com/christianchrisjo/hiring/internal/usecase"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type JobApplicationHandler struct {
	usecase *usecase.Usecase
}

func NewJobApplicationHandler(usecase *usecase.Usecase) *JobApplicationHandler {
	return &JobApplicationHandler{
		usecase: usecase,
	}
}

func (u *JobApplicationHandler) createJobApplication(w http.ResponseWriter, r *http.Request) {
	createRequest := models.CreateJobApplicationRequest{}
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &createRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "Invalid create user request")
		return
	}

	claimToken, err := extractBearerToken(r)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// get user id from token
	user, err := u.usecase.GetUserByID(claimToken.UserID.String())
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "invalid token credential")
		return
	}

	if user.UserID.String() != createRequest.UserID.String() {
		WriteWithResponse(w, http.StatusBadRequest, "unable to create job application for other user")
		return
	}

	job, err := u.usecase.CreateJobApplication(createRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusCreated, job)
}

func (u *JobApplicationHandler) getJobApplicationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	jobApplication, err := u.usecase.GetJobApplicationByID(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	claimToken, err := extractBearerToken(r)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if claimToken.Type == string(models.UserTypeEmployer) {
		user, err := u.usecase.GetUserByID(claimToken.UserID.String())
		if err != nil {
			WriteWithResponse(w, http.StatusBadRequest, "invalid token credential")
			return
		}
		job, err := u.usecase.GetJobByID(jobApplication.JobID.String())
		if err != nil {
			WriteWithResponse(w, http.StatusBadRequest, "invalid job id")
			return
		}
		if job.CompanyName != user.Name {
			WriteWithResponse(w, http.StatusBadRequest, "unable to see job application from other company")
			return
		}
	}
	if claimToken.Type == string(models.UserTypeEmployee) {
		user, err := u.usecase.GetUserByID(claimToken.UserID.String())
		if err != nil {
			WriteWithResponse(w, http.StatusBadRequest, "invalid token credential")
			return
		}
		if user.UserID.String() != jobApplication.UserID.String() {
			WriteWithResponse(w, http.StatusBadRequest, "unable to see other user job application")
			return
		}
	}

	WriteWithResponse(w, http.StatusOK, jobApplication)
}

func (u *JobApplicationHandler) getJobApplicationByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	claimToken, err := extractBearerToken(r)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// get user id from token
	user, err := u.usecase.GetUserByID(claimToken.UserID.String())
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "invalid token credential")
		return
	}

	if user.UserID.String() != id {
		WriteWithResponse(w, http.StatusBadRequest, "unable to see other user job application")
		return
	}

	jobApplications, err := u.usecase.GetJobApplicationByUserID(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, jobApplications)
}

func (u *JobApplicationHandler) getJobApplicationByJobID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	claimToken, err := extractBearerToken(r)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := u.usecase.GetUserByID(claimToken.UserID.String())
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "invalid token credential")
		return
	}
	job, err := u.usecase.GetJobByID(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "invalid job id")
		return
	}
	if job.CompanyName != user.Name {
		WriteWithResponse(w, http.StatusBadRequest, "unable to see job application from other company")
		return
	}

	jobApplications, err := u.usecase.GetJobApplicationsByJobID(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, jobApplications)
}

func (u *JobApplicationHandler) updateJobApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	updateRequest := models.UpdateJobApplicationRequest{}
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &updateRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "Invalid update job application request")
		return
	}
	updateRequest.ID, err = uuid.Parse(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "Invalid job application id")
		return
	}

	claimToken, err := extractBearerToken(r)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := u.usecase.GetUserByID(claimToken.UserID.String())
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "invalid token credential")
		return
	}
	jobApplication, err := u.usecase.GetJobApplicationByID(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "invalid job application id")
		return
	}
	job, err := u.usecase.GetJobByID(jobApplication.JobID.String())
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "invalid job id")
		return
	}
	if job.CompanyName != user.Name {
		WriteWithResponse(w, http.StatusBadRequest, "unable to see job application from other company")
		return
	}

	jobApplication, err = u.usecase.UpdateJobApplication(updateRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, jobApplication)
}
