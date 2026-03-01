package validators

import (
	"pmo/internal/pkg/response"
	"pmo/services/hr-service/features/employee/shared/dto"

	"strings"

	"github.com/go-playground/validator/v10"
)

type EmployeeValidator struct {
	validate *validator.Validate
}

func NewEmployeeValidator(v *validator.Validate) *EmployeeValidator {
	return &EmployeeValidator{
		validate: v,
	}
}
func (v *EmployeeValidator) ValidateCreateEmployee(req *dto.CreateEmployeeRequest) []response.Error {
	var errors []response.Error
	err := v.validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, response.Error{
				Field:   strings.ToLower(err.Field()),
				Code:    "VALIDATION_ERROR",
				Message: getValidationMessage(err),
			})
		}
	}
	// Validasi tambahan untuk photos
	if len(req.Photos) > 5 {
		errors = append(errors, response.Error{
			Field:   "photos",
			Code:    "VALIDATION_ERROR",
			Message: "Maximum 5 photos allowed",
		})
	}
	return errors
}
func (v *EmployeeValidator) ValidateUpdateEmployee(req *dto.UpdateEmployeeRequest) []response.Error {
	var errors []response.Error
	err := v.validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, response.Error{
				Field:   strings.ToLower(err.Field()),
				Code:    "VALIDATION_ERROR",
				Message: getValidationMessage(err),
			})
		}
	}
	return errors
}
func (v *EmployeeValidator) ValidateGetEmployeeByID(req *dto.GetEmployeeByIDRequest) []response.Error {
	var errors []response.Error
	err := v.validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, response.Error{
				Field: strings.ToLower(err.Field()), Code: "VALIDATION_ERROR",
				Message: getValidationMessage(err),
			})
		}
	}
	return errors
}
func (v *EmployeeValidator) ValidateGetAllEmployees(req *dto.GetAllEmployeesRequest) []response.Error {
	var errors []response.Error
	req.SetDefaults()
	err := v.validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, response.Error{
				Field:   strings.ToLower(err.Field()),
				Code:    "VALIDATION_ERROR",
				Message: getValidationMessage(err),
			})
		}
	}
	return errors
}
func getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too small"
	case "max":
		return "Value is too large"
	case "datetime":
		return "Invalid date format, use YYYY-MM-DD"
	default:
		return "Invalid value"
	}
}
func getPaginationValidationMessage(err validator.FieldError) string {
	switch err.Field() {
	case "Page":
		return "Page must be at least 1"
	case "PageSize":
		return "Page size must be between 1 and 100"
	case "SortBy":
		return "Sort by must be one of: employee_id, hire_date, department_id"
	case "SortDir":
		return "Sort direction must be either 'asc' or 'desc'"
	case "DepartmentName":
		return "First name must be between 1 and 30 characters"
	case "LocationID":
		return "Department ID must be a positive number"
	default:
		return getValidationMessage(err)
	}
}
