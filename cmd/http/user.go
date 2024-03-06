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

type UserHandler struct {
	usecase *usecase.Usecase
}

func NewUserHandler(usecase *usecase.Usecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (u *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	createRequest := models.CreateUserRequest{}
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &createRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "Invalid create user request")
		return
	}

	user, err := u.usecase.CreateUser(createRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusCreated, user)
}

func (u *UserHandler) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	user, err := u.usecase.GetUserByEmail(email)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, user)
}

func (u *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	claimToken, err := extractBearerToken(r)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	updateRequest := models.UpdateUserRequest{}
	reqBody, _ := io.ReadAll(r.Body)
	err = json.Unmarshal(reqBody, &updateRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "Invalid update user request")
		return
	}

	updateRequest.UserID, err = uuid.Parse(id)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	if claimToken.UserID.String() != updateRequest.UserID.String() {
		WriteWithResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	user, err := u.usecase.UpdateUser(updateRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteWithResponse(w, http.StatusOK, user)
}

func (u *UserHandler) signInWithEmail(w http.ResponseWriter, r *http.Request) {
	signInRequest := models.SignInByEmailRequest{}
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &signInRequest)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, "Invalid create user request")
		return
	}

	err = signInRequest.Validate()
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := u.usecase.SignInByEmail(signInRequest.Email, signInRequest.Password)
	if err != nil {
		WriteWithResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteWithResponse(w, http.StatusOK, token)
}
