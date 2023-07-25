package request

type EmployeRequest struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary   int    `json:"salary"`
}
