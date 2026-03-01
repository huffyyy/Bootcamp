package repository

import (
	"context"
	"pmo/services/hr-service/internal/domain/models"

	"gorm.io/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

// Employee operations
func (r *employeeRepository) Create(ctx context.Context, emp *models.EmployeeExt) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create employee

		if len(emp.Photos) > 0 {
			for i := range emp.Photos {
				emp.Photos[i].EmployeeID = emp.EmployeeID
			}
			if err := tx.Create(&emp.Photos).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
func (r *employeeRepository) GetByID(ctx context.Context, id int32) (*models.Employee, error) {
	var emp models.Employee
	err := r.db.WithContext(ctx).
		Preload("Photos").
		Where("employee_id = ?", id).
		First(&emp).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &emp, err
}
func (r *employeeRepository) GetByEmail(ctx context.Context, email string) (*models.Employee, error) {
	var emp models.Employee
	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&emp).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &emp, err
}
func (r *employeeRepository) Update(ctx context.Context, emp *models.Employee) error {
	return r.db.WithContext(ctx).Save(emp).Error
}
func (r *employeeRepository) Delete(ctx context.Context, id int32) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Delete photos first
		if err := tx.Where("employee_id = ?",
			id).Delete(&models.EmployeePhoto{}).Error; err != nil {
			return err
		}
		// Delete employee
		return tx.Delete(&models.Employee{}, "employee_id = ?", id).Error
	})
}
func (r *employeeRepository) FindAll(ctx context.Context, params *FindAllParams) ([]models.Employee, int64, error) {
	var employees []models.Employee
	var total int64
	query := r.db.WithContext(ctx).Model(&models.Employee{})
	// Apply filters
	if params.DepartmentID != nil {
		query = query.Where("department_id = ?", *params.DepartmentID)
	}
	if params.JobID != nil {
		query = query.Where("job_id = ?", *params.JobID)
	}
	if params.Search != "" {
		search := "%" + params.Search + "%"
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?",
			search, search, search)
	}
	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// Get paginated results with photos
	err := query.
		Preload("Photos").
		Order(params.GetOrder()).
		Offset(params.GetOffset()).
		Limit(params.GetLimit()).
		Find(&employees).Error
	return employees, total, err
}

// Photo operations
func (r *employeeRepository) AddPhotos(ctx context.Context, photos []models.EmployeePhoto) error {
	if len(photos) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Create(&photos).Error
}
func (r *employeeRepository) GetPhotos(ctx context.Context, employeeID int32) ([]models.EmployeePhoto, error) {
	var photos []models.EmployeePhoto
	err := r.db.WithContext(ctx).
		Where("employee_id = ?", employeeID).
		Order("is_primary DESC, epho_id ASC").
		Find(&photos).Error
	return photos, err
}
func (r *employeeRepository) DeletePhoto(ctx context.Context, photoID int32) error {
	return r.db.WithContext(ctx).
		Delete(&models.EmployeePhoto{}, "epho_id = ?", photoID).Error
}
func (r *employeeRepository) SetPrimaryPhoto(ctx context.Context, employeeID, photoID int32) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Reset all photos to non-primary
		if err := tx.Model(&models.EmployeePhoto{}).
			Where("employee_id = ?", employeeID).
			Update("is_primary", false).Error; err != nil {
			return err
		}
		// Set selected photo as primary
		return tx.Model(&models.EmployeePhoto{}).
			Where("epho_id = ? AND employee_id = ?", photoID, employeeID).
			Update("is_primary", true).Error
	})
}
