package repository

import "errors"

var (
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrEmployeeNotFound  = errors.New("employee not found")
	ErrPhotoNotFound     = errors.New("photo not found")
	ErrInvalidEmployeeID = errors.New("invalid employee ID")
)
