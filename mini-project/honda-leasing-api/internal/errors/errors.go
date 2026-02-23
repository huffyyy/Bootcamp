package errs

import "errors"

var (
	ErrInvalidInput      = errors.New("invalid input")
	ErrInvalidPagination = errors.New("invalid pagination parameters")
	ErrInvalidSort       = errors.New("invalid sort parameters")
	ErrInvalidSearch     = errors.New("search name cannot be empty")
	ErrInvalidEmail      = errors.New("email already exist")
	ErrInvalidPassword   = errors.New("invalid password")

	ErrCreateUser = errors.New("error when create user")
	ErrAssignRole = errors.New("error when assigned user role")
	ErrTrxUser    = errors.New("transaction errors")

	ErrMotorNotFound    = errors.New("motor not found")
	ErrProductNotFound  = errors.New("product not found")
	ErrCustomerNotFound = errors.New("customer not found")
	ErrInvalidDP        = errors.New("invalid dp")
	ErrInvalidTenor     = errors.New("invalid tenor")
	ErrCreateContract   = errors.New("failed to create contract")

	ErrContractNotFound = errors.New("contract not found")
	ErrTaskNotFound     = errors.New("task not found")
)
