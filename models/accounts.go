package models

import "time"

type Account struct {
	ID           string    `json:"id"`
	EmployeeCode string    `json:"employee_code"`
	Token        string    `json:"token"`
	EmployeeName string    `json:"employee_name"`
	LastAccessed time.Time `json:"last_accessed"`
}
