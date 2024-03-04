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

	job, err := u.usecase.UpdateJob(updateRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, job)
}

func (u *JobHandler) deleteJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := u.usecase.DeleteJob(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, "Job deleted")
}
