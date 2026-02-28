package repository

import "errors"

var (
	ErrDuplicateName       = errors.New("department name already exists")
	ErrDepartmentNotFound  = errors.New("department not found")
	ErrInvalidDepartmentID = errors.New("invalid department ID")
)
