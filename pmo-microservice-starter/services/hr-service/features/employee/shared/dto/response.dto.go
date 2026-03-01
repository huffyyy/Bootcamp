package dto

import (
	"pmo/services/hr-service/internal/domain/models"
	"time"
)

type EmployeePhotoResponse struct {
	EphoID    int32   `json:"epho_id"`
	FileName  *string `json:"file_name"`
	FileSize  *int64  `json:"file_size"`
	FileType  *string `json:"file_type"`
	FileURL   *string `json:"file_url"`
	IsPrimary *bool   `json:"is_primary"`
}

type EmployeeResponse struct {
	EmployeeID   int32                   `json:"employee_id"`
	FirstName    *string                 `json:"first_name"`
	LastName     string                  `json:"last_name"`
	Email        string                  `json:"email"`
	PhoneNumber  *string                 `json:"phone_number"`
	HireDate     time.Time               `json:"hire_date"`
	JobID        int32                   `json:"job_id"`
	Salary       float64                 `json:"salary"`
	ManagerID    *int32                  `json:"manager_id"`
	DepartmentID *int32                  `json:"department_id"`
	Photos       []EmployeePhotoResponse `json:"photos,omitempty"`
}

type EmployeeListResponse struct {
	Employees []EmployeeResponse `json:"employees"`
	Total     int64              `json:"total"`
}

func ToEmployeePhotoResponse(photo *models.EmployeePhoto) EmployeePhotoResponse {
	if photo == nil {
		return EmployeePhotoResponse{}
	}
	return EmployeePhotoResponse{
		EphoID:    photo.EphoID,
		FileName:  photo.FileName,
		FileSize:  photo.FileSize,
		FileType:  photo.FileType,
		FileURL:   photo.FileURL,
		IsPrimary: photo.IsPrimary,
	}
}

func ToEmployeePhotoResponses(photos []models.EmployeePhoto) []EmployeePhotoResponse {
	responses := make([]EmployeePhotoResponse, len(photos))
	for i, photo := range photos {
		responses[i] = ToEmployeePhotoResponse(&photo)
	}
	return responses
}
func ToEmployeeResponse(emp *models.EmployeeExt) *EmployeeResponse {
	if emp == nil {
		return nil
	}
	return &EmployeeResponse{
		EmployeeID:   emp.EmployeeID,
		FirstName:    emp.FirstName,
		LastName:     emp.LastName,
		Email:        emp.Email,
		PhoneNumber:  emp.PhoneNumber,
		HireDate:     emp.HireDate,
		JobID:        emp.JobID,
		Salary:       emp.Salary,
		ManagerID:    emp.ManagerID,
		DepartmentID: emp.DepartmentID,
		Photos:       ToEmployeePhotoResponses(emp.Photos),
	}
}

func ToEmployeeResponses(emps []models.EmployeeExt) []EmployeeResponse {
	responses := make([]EmployeeResponse, len(emps))
	for i, emp := range emps {
		response := ToEmployeeResponse(&emp)
		if response != nil {
			responses[i] = *response
		}
	}
	return responses
}
func ToEmployeeListResponse(emps []models.EmployeeExt, total int64) *EmployeeListResponse {
	return &EmployeeListResponse{
		Employees: ToEmployeeResponses(emps),
		Total:     total,
	}
}
