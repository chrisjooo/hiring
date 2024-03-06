simple job portal API

## How to Run
go to `./cmd` where `main.go` file exist. simply run `go run main.go`. The app itself will be loaded in `localhost:10000`

## API Documentation
https://docs.google.com/document/d/1h9_mLZ_xJkqLH5CFkUvtoSawVxuVJ8fBquPITVQ3iBM/edit?usp=sharing

## Bearer token
to get the token, we need to have user via `[POST] /user`. After user is created, proceed to sign in using `[POST] /user/signin`. the sign in endpoint will return string. <br>
This is the bearer token which you will need to attach it in every request metadata `authorization: bearer {token}` <br>
Each bearer token is unique and save some of user information such as type. some endpoint will need bearer token with certain user type

## API list:
#### Endpoints without bearer token
```
	"[POST] /user : Create a new user"
	"[POST] /user/signin : Sign in"
	"[GET] /job/list : Get all jobs list"
	"[GET] /job/{id} : Get job detail by ID"
```
#### Endpoints with general user bearer token
```
	"[GET] /user/{email} : Get user by email"
	"[PUT] /user/{id} : Update user by user ID"
	"[GET] /job/application/{id} : Get user job application detail by job application ID"
```
#### Endpoints with employer bearer token
```
	"[POST] /job : Create a new job"
	"[PUT] /job/{id} : Update a job"
	"[DELETE] /job/{id} : Delete a job by job id"
	"[GET] /job/application/job/{id} : get all applications by job id"
	"[PUT] /job/application/{id} : update a job application by id"
```
#### Endpoints with employee bearer token
```
	[GET] /job/application/user/{id} : Get all job applications by user id"
	[POST] /job/application : Create a new job application"
```
