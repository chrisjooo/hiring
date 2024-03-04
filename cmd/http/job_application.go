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

	job, err := u.usecase.CreateJobApplication(createRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusCreated, job)
}

func (u *JobApplicationHandler) getJobApplication(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	jobID := r.URL.Query().Get("job_id")
	jobApplicationID := r.URL.Query().Get("id")

	if jobApplicationID != "" {
		jobApplication, err := u.usecase.GetJobApplicationByID(jobApplicationID)
		if err != nil {
			WriteWithResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		WriteWithResponse(w, http.StatusOK, jobApplication)
	}
	if userID != "" && jobID != "" {
		jobApplication, err := u.usecase.GetJobApplicationByJobIDAndUserID(jobID, userID)
		if err != nil {
			WriteWithResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		WriteWithResponse(w, http.StatusOK, jobApplication)
	}
	if userID != "" {
		jobApplications, err := u.usecase.GetJobApplicationByUserID(userID)
		if err != nil {
			WriteWithResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		WriteWithResponse(w, http.StatusOK, jobApplications)
	}
	if jobID != "" {
		jobApplications, err := u.usecase.GetJobApplicationsByJobID(jobID)
		if err != nil {
			WriteWithResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		WriteWithResponse(w, http.StatusOK, jobApplications)
	}
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
	jobApplication, err := u.usecase.UpdateJobApplication(updateRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, jobApplication)
}
