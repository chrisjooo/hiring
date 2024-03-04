package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/christianchrisjo/hiring/internal/usecase"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoints without bearer token")
	fmt.Fprintln(w, "[GET] /user/{email} : Get a user by email")
	fmt.Fprintln(w, "[POST] /user/create : create a new user")
	fmt.Fprintln(w, "[GET] /job/list : List all jobs")
	fmt.Fprintln(w, "[GET] /job/{id} : Get a job by id")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "Endpoints with employee bearer token")
	fmt.Fprintln(w, "[PUT] /user/{id} : create a new user")
	fmt.Fprintln(w, "")

	// TODO add employer bearer token authentication
	fmt.Fprintln(w, "Endpoints with employer bearer token")
	fmt.Fprintln(w, "[POST] /job/create : Create a new job")
	fmt.Fprintln(w, "[PUT] /job/{id} : update a job by id")
	fmt.Fprintln(w, "[DELETE] /job/{id} : delete a job by id")
}

func HandleRequests(handler *Handlers) {
	myRouter := mux.NewRouter().StrictSlash(true)

	// endpoint documentation
	myRouter.HandleFunc("/", homePage).Methods("GET")

	// user endpoints
	myRouter.HandleFunc("/user/{email}", handler.userHandler.getUserByEmail).Methods("GET")
	myRouter.HandleFunc("/user", handler.userHandler.createUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", handler.userHandler.updateUser).Methods("PUT")

	// job endpoints
	myRouter.HandleFunc("/job/list", handler.jobHandler.getAllJobs).Methods("GET")
	myRouter.HandleFunc("/job/{id}", handler.jobHandler.getJobByID).Methods("GET")
	myRouter.HandleFunc("/job", handler.jobHandler.createJob).Methods("POST")
	myRouter.HandleFunc("/job/{id}", handler.jobHandler.updateJob).Methods("PUT")
	myRouter.HandleFunc("/job/{id}", handler.jobHandler.deleteJob).Methods("DELETE")

	// job application endpoints
	myRouter.HandleFunc("/job/application/", handler.jobApplicationHandler.getJobApplication).Methods("GET")
	myRouter.HandleFunc("/job/application/", handler.jobApplicationHandler.createJobApplication).Methods("POST")
	myRouter.HandleFunc("/job/application/{id}", handler.jobApplicationHandler.updateJobApplication).Methods("PUT")

	// CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"})

	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(originsOk, headersOk, methodsOk)(myRouter)))
}

type Handlers struct {
	userHandler           *UserHandler
	jobHandler            *JobHandler
	jobApplicationHandler *JobApplicationHandler
}

func NewHandlers(uc *usecase.Usecase) *Handlers {
	return &Handlers{
		userHandler:           NewUserHandler(uc),
		jobHandler:            NewJobHandler(uc),
		jobApplicationHandler: NewJobApplicationHandler(uc),
	}
}
