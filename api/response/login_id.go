package response

type LoginByIDResponse struct {
	EmployeeCode string `json:"employee_code"`
	EmployeeName string `json:"employee_name"`
}
