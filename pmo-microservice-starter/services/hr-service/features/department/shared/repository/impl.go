package repository

import (
	"context"
	"errors"
	"pmo/services/hr-service/internal/domain/models"

	"gorm.io/gorm"
)

type departmentRepository struct {
	db *gorm.DB
}

func (r *departmentRepository) Create(ctx context.Context, dept *models.Department) error {
	return r.db.WithContext(ctx).Create(dept).Error
}

func (r *departmentRepository) GetByID(ctx context.Context, id int32) (*models.Department, error) {
	var dept models.Department
	err := r.db.WithContext(ctx).Where("department_id = ?", id).First(&dept).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &dept, err
}

func (r *departmentRepository) GetByName(ctx context.Context, name string) ([]models.Department, error) {
	var departments []models.Department
	err := r.db.WithContext(ctx).Where("department_name ILIKE ?",
		"%"+name+"%").Find(&departments).Error
	return departments, err
}

func (r *departmentRepository) Update(ctx context.Context, dept *models.Department) error {
	return r.db.WithContext(ctx).Save(dept).Error
}

func (r *departmentRepository) Delete(ctx context.Context, id int32) error {
	return r.db.WithContext(ctx).Delete(&models.Department{}, "department_id = ?",
		id).Error
}

func (r *departmentRepository) GetAll(ctx context.Context) ([]models.Department,
	error) {
	var departments []models.Department
	err := r.db.WithContext(ctx).Find(&departments).Error
	return departments, err
}

func (r *departmentRepository) FindAll(ctx context.Context, params *FindAllParams) ([]models.Department, int64, error) {
	var departments []models.Department
	var total int64
	// Build query
	query := r.db.WithContext(ctx).Model(&models.Department{})
	// Apply filters
	if params.DepartmentName != "" {
		query = query.Where("department_name ILIKE ?",
			"%"+params.DepartmentName+"%")
	}
	if params.LocationID != nil {
		query = query.Where("location_id = ?", *params.LocationID)
	}
	// Get total count before pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// Apply pagination and sorting
	err := query.
		Order(params.GetOrder()).
		Offset(params.GetOffset()).
		Limit(params.GetLimit()).
		Find(&departments).Error
	if err != nil {
		return nil, 0, err
	}
	return departments, total, nil
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{
		db: db,
	}
}
