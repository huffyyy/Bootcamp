package employee

import "errors"

var (
	ErrInvalidId 		= errors.New("ID mus be positive")
	ErrEmptyFirstName 	= errors.New("first name cannot be empty")
	ErrEmptyLastName 	= errors.New("last name cannot be empty")
	ErrInvalidSalaryMin = errors.New("Salary must be positive")
	ErrInvalidSalaryMax = errors.New("salary must be less than Rp.1jt")
	ErrFutureHireDate 	= errors.New("hire date cannoto be in the future")
)