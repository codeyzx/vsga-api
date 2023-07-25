## VSGA-API

VSGA API is a RESTful API built with Go (Golang) and Postgres. It is the backend for the [VSGA](https://github.com/codeyzx/vsga) Mobile Application.

## Getting Started

1. Clone the repository: `git clone https://github.com/codeyzx/vsga-api.git`
2. Navigate to the project directory: `cd vsga-api`
3. Install the dependencies: `go get -u`
4. Create database: `vsga-api`
5. Run the application: `go run main.go`

> notes: database will be migrated automatically

## API Documentation

| METHOD   | ROUTE          | FUNCTIONALITY                           | ACCESS      |
| -------- | -------------- | --------------------------------------- | ----------- |
| _GET_    | `/`            | _Redirects to API documentation_        | _All users_ |
| _GET_    | `/docs/*`      | _Serves the API documentation page_     | _All users_ |
| _GET_    | `/employe`     | _Gets a list of all employee_           | _All users_ |
| _GET_    | `/employe/:id` | _Gets information on a single employee_ | _All users_ |
| _POST_   | `/employe`     | _Creates a new employee_                | _All users_ |
| _PUT_    | `/employe/:id` | _Update an employee_                    | _All users_ |
| _DELETE_ | `/employe/:id` | _Delete an employee_                    | _All users_ |
