package dto

import (
	"mime/multipart"
)

type CreateEmployeeRequest struct {
	FirstName    *string `form:"first_name" validate:"omitempty,max=20"`
	LastName     string  `form:"last_name" validate:"required,max=25"`
	Email        string  `form:"email" validate:"required,email,max=100"`
	PhoneNumber  *string `form:"phone_number" validate:"omitempty,max=20"`
	HireDate     string  `form:"hire_date" validate:"required,datetime=2006-01-02"`
	JobID        int32   `form:"job_id" validate:"required,min=1"`
	Salary       float64 `form:"salary" validate:"required,min=0"`
	ManagerID    *int32  `form:"manager_id" validate:"omitempty,min=1"`
	DepartmentID *int32  `form:"department_id" validate:"omitempty,min=1"`
	// Upload photos
	Photos []*multipart.FileHeader `form:"photos"` // max 5MB per photo :validate:"omitempty,dive,max=5120
}
type UpdateEmployeeRequest struct {
	FirstName    *string  `form:"first_name" validate:"omitempty,max=20"`
	LastName     *string  `form:"last_name" validate:"omitempty,max=25"`
	Email        *string  `form:"email" validate:"omitempty,email,max=100"`
	PhoneNumber  *string  `form:"phone_number" validate:"omitempty,max=20"`
	HireDate     *string  `form:"hire_date" validate:"omitempty,datetime=2006-01-02"`
	JobID        *int32   `form:"job_id" validate:"omitempty,min=1"`
	Salary       *float64 `form:"salary" validate:"omitempty,min=0"`
	ManagerID    *int32   `form:"manager_id" validate:"omitempty,min=1"`
	DepartmentID *int32   `form:"department_id" validate:"omitempty,min=1"`
}
type GetEmployeeByIDRequest struct {
	EmployeeID int32 `uri:"id" validate:"required,min=1"`
}
type GetAllEmployeesRequest struct {
	Page         int    `form:"page" validate:"omitempty,min=1"`
	PageSize     int    `form:"page_size" validate:"omitempty,min=1,max=100"`
	SortBy       string `form:"sort_by" validate:"omitempty,oneof=employee_id first_name last_name email hire_date salary"`
	SortDir      string `form:"sort_dir" validate:"omitempty,oneof=asc desc"`
	DepartmentID *int32 `form:"department_id" validate:"omitempty,min=1"`
	JobID        *int32 `form:"job_id" validate:"omitempty,min=1"`
	Search       string `form:"search" validate:"omitempty,max=50"`
}

func (r *GetAllEmployeesRequest) SetDefaults() {
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.PageSize <= 0 {
		r.PageSize = 10
	}
	if r.SortBy == "" {
		r.SortBy = "employee_id"
	}
	if r.SortDir == "" {
		r.SortDir = "asc"
	}
}
