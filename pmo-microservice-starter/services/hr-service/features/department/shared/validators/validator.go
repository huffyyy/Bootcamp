package validators

import (
	"pmo/internal/pkg/response"
	"pmo/services/hr-service/features/department/shared/dto"
	"strings"

	"github.com/go-playground/validator/v10"
)

type DepartmentValidator struct {
	validate *validator.Validate
}

func NewDepartmentValidator(v *validator.Validate) *DepartmentValidator {
	return &DepartmentValidator{
		validate: v,
	}
}

func (v *DepartmentValidator) ValidateCreateDepartment(req *dto.CreateDepartmentRequest) []response.Error {
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

func (v *DepartmentValidator) ValidateUpdateDepartment(req *dto.UpdateDepartmentRequest) []response.Error {
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

func (v *DepartmentValidator) ValidateGetDepartmentByID(req *dto.GetDepartmentByIDRequest) []response.Error {
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

func (v *DepartmentValidator) ValidateGetDepartmentByName(req *dto.GetDepartmentByNameRequest) []response.Error {
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

func (v *DepartmentValidator) ValidateDeleteDepartment(req *dto.DeleteDepartmentRequest) []response.Error {
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

func (v *DepartmentValidator) ValidateGetAllDepartments(req *dto.GetAllDepartmentsRequest) []response.Error {
	var errors []response.Error
	// Set default values first
	req.SetDefaults()
	err := v.validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, response.Error{
				Field:   strings.ToLower(err.Field()),
				Code:    "VALIDATION_ERROR",
				Message: getPaginationValidationMessage(err),
			})
		}
	}
	return errors
}

func getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
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
		return "Sort by must be one of: department_id, department_name, location_id"
	case "SortDir":
		return "Sort direction must be either 'asc' or 'desc'"
	case "DepartmentName":
		return "Department name must be between 1 and 30 characters"
	case "LocationID":
		return "Location ID must be a positive number"
	default:
		return getValidationMessage(err)
	}
}
