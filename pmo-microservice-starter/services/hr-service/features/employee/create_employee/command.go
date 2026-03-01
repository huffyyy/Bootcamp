package createemployee

import (
	"mime/multipart"
	"pmo/services/hr-service/internal/domain/models"
	"time"
)

type CreateEmployeeCommand struct {
	FirstName    *string
	LastName     string
	Email        string
	PhoneNumber  *string
	HireDate     time.Time
	JobID        int32
	Salary       float64
	ManagerID    *int32
	DepartmentID *int32
	Photos       []*multipart.FileHeader
}

func NewCreateEmployeeCommand(
	firstName *string,
	lastName string,
	email string,
	phoneNumber *string,
	hireDate time.Time,
	jobID int32,
	salary float64,
	managerID *int32,
	departmentID *int32,
	photos []*multipart.FileHeader,
) CreateEmployeeCommand {
	return CreateEmployeeCommand{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PhoneNumber:  phoneNumber,
		HireDate:     hireDate,
		JobID:        jobID,
		Salary:       salary,
		ManagerID:    managerID,
		DepartmentID: departmentID,
		Photos:       photos,
	}
}

func (c *CreateEmployeeCommand) ToModel() *models.EmployeeExt {
	return &models.EmployeeExt{
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		Email:        c.Email,
		PhoneNumber:  c.PhoneNumber,
		HireDate:     c.HireDate,
		JobID:        c.JobID,
		Salary:       c.Salary,
		ManagerID:    c.ManagerID,
		DepartmentID: c.DepartmentID,
		Photos:       []models.EmployeePhoto{},
	}
}
