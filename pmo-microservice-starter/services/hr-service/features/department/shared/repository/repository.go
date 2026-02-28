package repository

import (
	"context"
	"pmo/services/hr-service/internal/domain/models"
)

type DepartmentRepository interface {
	Create(ctx context.Context, dept *models.Department) error
	GetByID(ctx context.Context, id int32) (*models.Department, error)
	GetByName(ctx context.Context, name string) ([]models.Department, error)
	Update(ctx context.Context, dept *models.Department) error
	Delete(ctx context.Context, id int32) error
	GetAll(ctx context.Context) ([]models.Department, error)
	// New methods for pagination
	FindAll(ctx context.Context, params *FindAllParams) ([]models.Department, int64, error)
}

type FindAllParams struct {
	Page     int
	PageSize int
	SortBy   string
	SortDir  string
	// Filters
	DepartmentName string
	LocationID     *int32
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
