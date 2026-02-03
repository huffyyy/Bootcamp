package dto

import "time"

// CreateEmployeeRequest for creating a employee
type CreateEmployeeRequest struct {
	FirstName    string  `json:"first_name" validate:"omitempty,min=2,max=20"`
	LastName     string  `json:"last_name" validate:"required,min=2,max=25"`
	Email        string  `json:"email" validate:"required,email,max=100"`
	PhoneNumber  string  `json:"phone_number" validate:"omitempty,min=10,max=20"`
	HireDate     string  `json:"hire_date" validate:"required"`
	JobID        int32   `json:"job_id" validate:"required,min=1"`
	Salary       float64 `json:"salary" validate:"required,min=0"`
	ManagerID    int32   `json:"manager_id" validate:"omitempty,min=1"`             
	DepartmentID int32   `json:"department_id" validate:"omitempty,min=1"`
}

// UpdateEmployeeRequest for partial update
type UpdateEmployeeRequest struct {
	FirstName    *string  `json:"first_name,omitempty" validate:"omitempty,min=2,max=20"`
	LastName     *string  `json:"last_name,omitempty" validate:"omitempty,min=2,max=25"`
	Email        *string  `json:"email,omitempty" validate:"omitempty,email,max=100"`
	PhoneNumber  *string  `json:"phone_number,omitempty" validate:"omitempty,min=10,max=20"`
	HireDate     *string  `json:"hire_date,omitempty"`
	JobID        *int32   `json:"job_id,omitempty" validate:"omitempty,min=1"`
	Salary       *float64 `json:"salary,omitempty" validate:"omitempty,min=0"`
	ManagerID    *int32   `json:"manager_id,omitempty" validate:"omitempty,min=1"`
	DepartmentID *int32   `json:"department_id,omitempty" validate:"omitempty,min=1"`
}

// EmployeeResponse for output (hide internal fields if needed)
type EmployeeResponse struct {
	ID           int32     `json:"id"`
	FirstName    *string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PhoneNumber  *string    `json:"phone_number"`
	HireDate     time.Time `json:"hire_date"`
	JobID        int32     `json:"job_id"`
	Salary       float64   `json:"salary"`
	ManagerID    *int32    `json:"manager_id,omitempty"`
	DepartmentID *int32   `json:"department_id,omitempty"`
}