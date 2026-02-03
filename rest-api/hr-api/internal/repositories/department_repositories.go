package repositories

import (
	"context"

	"github.com/codeid/hr-api/internal/domain/models"
	"github.com/codeid/hr-api/internal/domain/query"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	FindAll(ctx context.Context) ([]*models.Department, error)
	FindByID(ctx context.Context, id uint) (*models.Department, error)
	Create(ctx context.Context, department *models.Department) error
	Update(ctx context.Context, department *models.Department) error
	Delete(ctx context.Context, id uint) error
	SearchByName(ctx context.Context, name string) ([]*models.Department, error)
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{Q: query.Use(db)}
}

type departmentRepository struct {
	Q *query.Query
}

// SearchByName implements DepartmentRepository.
func (r *departmentRepository) SearchByName(ctx context.Context, name string) ([]*models.Department, error) {
	departments, err := r.Q.Department.WithContext(ctx).
		Where(r.Q.Department.DepartmentName.Like("%" + name + "%")).Find()
	if err != nil {
		return nil, err
	}
	return departments, nil
}

// Create implements DepartmentRepository.
func (r *departmentRepository) Create(ctx context.Context, department *models.Department) error {
	return r.Q.Department.WithContext(ctx).Create(department)
}

// Delete implements DepartmentRepository.
func (r *departmentRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.Q.Department.WithContext(ctx).Where(r.Q.Department.DepartmentID.Eq(int32(id))).Delete(&models.Department{})
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements DepartmentRepository.
func (r *departmentRepository) FindAll(ctx context.Context) ([]*models.Department, error) {
	var departments []*models.Department
	departments, err := r.Q.Department.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	return departments, err
}

// FindByID implements DepartmentRepository.
func (r *departmentRepository) FindByID(ctx context.Context, id uint) (*models.Department, error) {
	department, err := r.Q.Department.WithContext(ctx).Where(r.Q.Department.DepartmentID.Eq(int32(id))).First()
	if err != nil {
		return nil, err
	}
	return department, nil
}

// Update implements DepartmentRepository.
func (r *departmentRepository) Update(ctx context.Context, department *models.Department) error {
	return r.Q.Department.WithContext(ctx).Save(department)
}