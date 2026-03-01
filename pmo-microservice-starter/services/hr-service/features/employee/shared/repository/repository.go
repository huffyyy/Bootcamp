package repository

import (
	"context"
	"pmo/services/hr-service/internal/domain/models"
)

type EmployeeRepository interface {
	Create(ctx context.Context, emp *models.EmployeeExt) error
	GetByID(ctx context.Context, id int32) (*models.Employee, error)
	GetByEmail(ctx context.Context, email string) (*models.Employee, error)
	Update(ctx context.Context, emp *models.Employee) error
	Delete(ctx context.Context, id int32) error
	FindAll(ctx context.Context, params *FindAllParams) ([]models.Employee, int64, error)
	// Photo operations
	AddPhotos(ctx context.Context, photos []models.EmployeePhoto) error
	GetPhotos(ctx context.Context, employeeID int32) ([]models.EmployeePhoto, error)
	DeletePhoto(ctx context.Context, photoID int32) error
	SetPrimaryPhoto(ctx context.Context, employeeID, photoID int32) error
}

type FindAllParams struct {
	Page         int
	PageSize     int
	SortBy       string
	SortDir      string
	DepartmentID *int32
	JobID        *int32
	Search       string
}

func (p *FindAllParams) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *FindAllParams) GetLimit() int {
	return p.PageSize
}

func (p *FindAllParams) GetOrder() string {
	return p.SortBy + " " + p.SortDir
}
