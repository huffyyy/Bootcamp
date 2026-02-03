package services

import (
	"context"
	"time"

	"github.com/codeid/hr-api/internal/domain/models"
	"github.com/codeid/hr-api/internal/dto"
	errs "github.com/codeid/hr-api/internal/errors"
	"github.com/codeid/hr-api/internal/repositories"
	"github.com/go-playground/validator/v10"
)

type EmployeeServiceInterface interface {
	Create(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error)
	FindByID(ctx context.Context, id uint) (*dto.EmployeeResponse, error)
	GetAll(ctx context.Context) ([]dto.EmployeeResponse, error)
	Update(ctx context.Context, id uint, req *dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error)
	Delete(ctx context.Context, id uint) error
	SearchByName(ctx context.Context, name string) ([]dto.EmployeeResponse, error)
}

type employeeService struct {
	repo     repositories.EmployeeRepository
	validate *validator.Validate
}

// helper to safely dereference optional fields
func strPtrToString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func uintPtrToUint(p *uint) uint {
	if p == nil {
		return 0
	}
	return *p
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeServiceInterface {
	return &employeeService{
		repo:     repo,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

// Create implements [EmployeeServiceInterface].
func (s *employeeService) Create(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, errs.ErrInvalidInput
	}

	hireDate, err := time.Parse("2006-01-02", req.HireDate)
	if err != nil {
		return nil, errs.ErrInvalidInput
	}

	employee := &models.Employee{
		LastName: req.LastName,
		Email:    req.Email,
		HireDate: hireDate,
		JobID:    req.JobID,
		Salary:   req.Salary,
	}

	if req.FirstName != "" {
		employee.FirstName = &req.FirstName
	}
	if req.PhoneNumber != "" {
		employee.PhoneNumber = &req.PhoneNumber
	}
	if req.ManagerID != 0 {
		employee.ManagerID = &req.ManagerID
	}
	if req.DepartmentID != 0 {
		employee.DepartmentID = &req.DepartmentID
	}

	if err := s.repo.Create(ctx, employee); err != nil {
		return nil, err
	}

	return &dto.EmployeeResponse{
		ID:           employee.EmployeeID,
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		Email:        employee.Email,
		PhoneNumber:  employee.PhoneNumber,
		HireDate:     employee.HireDate,
		JobID:        employee.JobID,
		Salary:       employee.Salary,
		ManagerID:    employee.ManagerID,
		DepartmentID: employee.DepartmentID,
	}, nil
}

// FindByID implements [EmployeeServiceInterface].
func (s *employeeService) FindByID(ctx context.Context, id uint) (*dto.EmployeeResponse, error) {
	emp, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.EmployeeResponse{
		ID:           emp.EmployeeID,
		FirstName:    emp.FirstName,
		LastName:     emp.LastName,
		Email:        emp.Email,
		PhoneNumber:  emp.PhoneNumber,
		HireDate:     emp.HireDate,
		JobID:        emp.JobID,
		Salary:       emp.Salary,
		ManagerID:    emp.ManagerID,
		DepartmentID: emp.DepartmentID,
	}, nil
}

// GetAll implements [EmployeeServiceInterface].
func (s *employeeService) GetAll(ctx context.Context) ([]dto.EmployeeResponse, error) {
	emps, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.EmployeeResponse
	for _, emp := range emps {
		responses = append(responses, dto.EmployeeResponse{
			ID:           emp.EmployeeID,
			FirstName:    emp.FirstName,
			LastName:     emp.LastName,
			Email:        emp.Email,
			PhoneNumber:  emp.PhoneNumber,
			HireDate:     emp.HireDate,
			JobID:        emp.JobID,
			Salary:       emp.Salary,
			ManagerID:    emp.ManagerID,
			DepartmentID: emp.DepartmentID,
		})
	}
	return responses, nil
}

// Update implements [EmployeeServiceInterface].
func (s *employeeService) Update(ctx context.Context, id uint, req *dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, errs.ErrInvalidInput
	}

	emp, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Map DTO to model (partial update) 
	if req.FirstName != nil {
		emp.FirstName = req.FirstName
	}
	if req.LastName != nil {
		emp.LastName = *req.LastName
	}
	if req.Email != nil {
		emp.Email = *req.Email
	}
	if req.PhoneNumber != nil {
		emp.PhoneNumber = req.PhoneNumber
	}
	if req.HireDate != nil {
		hireDate, err := time.Parse("2006-01-02", *req.HireDate)
		if err != nil {
			return nil, errs.ErrInvalidInput
		}
		emp.HireDate = hireDate
	}
	if req.JobID != nil {
		emp.JobID = *req.JobID
	}
	if req.Salary != nil {
		emp.Salary = *req.Salary
	}
	if req.ManagerID != nil {
		emp.ManagerID = req.ManagerID
	}
	if req.DepartmentID != nil {
		emp.DepartmentID = req.DepartmentID
	}

	if err := s.repo.Update(ctx, emp); err != nil {
		return nil, err
	}

	return &dto.EmployeeResponse{
		ID:           emp.EmployeeID,
		FirstName:    emp.FirstName,
		LastName:     emp.LastName,
		Email:        emp.Email,
		PhoneNumber:  emp.PhoneNumber,
		HireDate:     emp.HireDate,
		JobID:        emp.JobID,
		Salary:       emp.Salary,
		ManagerID:    emp.ManagerID,
		DepartmentID: emp.DepartmentID,
	}, nil
}


// Delete implements [EmployeeServiceInterface].
func (s *employeeService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// SearchByName implements [EmployeeServiceInterface].
func (s *employeeService) SearchByName(ctx context.Context, name string) ([]dto.EmployeeResponse, error) {
	if name == "" {
		return s.GetAll(ctx)
	}

	emps, err := s.repo.SearchByName(ctx, name)
	if err != nil {
		return nil, err
	}

	var responses []dto.EmployeeResponse
	for _, emp := range emps {
		responses = append(responses, dto.EmployeeResponse{
			ID:           emp.EmployeeID,
			FirstName:    emp.FirstName,
			LastName:     emp.LastName,
			Email:        emp.Email,
			PhoneNumber:  emp.PhoneNumber,
			HireDate:     emp.HireDate,
			JobID:        emp.JobID,
			Salary:       emp.Salary,
			ManagerID:    emp.ManagerID,
			DepartmentID: emp.DepartmentID,
		})
	}
	return responses, nil
}