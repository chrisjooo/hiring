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

type JobHandler struct {
	usecase *usecase.Usecase
}

func NewJobHandler(usecase *usecase.Usecase) *JobHandler {
	return &JobHandler{
		usecase: usecase,
	}
}

func (u *JobHandler) createJob(w http.ResponseWriter, r *http.Request) {
	createRequest := models.CreateJobRequest{}
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

	if user.Name != createRequest.CompanyName {
		WriteWithResponse(w, http.StatusBadRequest, "unable to create job for other company")
		return
	}

	job, err := u.usecase.CreateJob(createRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusCreated, job)
}

func (u *JobHandler) getJobByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	job, err := u.usecase.GetJobByID(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, job)
}

func (u *JobHandler) getAllJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := u.usecase.GetAllJobs()
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, jobs)
}

func (u *JobHandler) updateJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	updateRequest := models.UpdateJobRequest{}
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &updateRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "Invalid update job request")
		return
	}

	updateRequest.ID, err = uuid.Parse(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "Invalid job id")
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
	job, err := u.usecase.GetJobByID(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "invalid job id")
		return
	}

	if user.Name != job.CompanyName {
		WriteWithResponse(w, http.StatusBadRequest, "unable to update other company job")
		return
	}

	job, err = u.usecase.UpdateJob(updateRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, job)
}

func (u *JobHandler) deleteJob(w http.ResponseWriter, r *http.Request) {
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

	job, err := u.usecase.GetJobByID(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "invalid job id")
		return
	}

	if user.Name != job.CompanyName {
		WriteWithResponse(w, http.StatusBadRequest, "unable to update other company job")
		return
	}

	err = u.usecase.DeleteJob(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, "Job deleted")
}
