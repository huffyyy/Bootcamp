package repositories

import (
	"context"

	"github.com/codeid/hr-api/internal/domain/models"
	"github.com/codeid/hr-api/internal/domain/query"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll(ctx context.Context) ([]*models.Employee, error)
	FindByID(ctx context.Context, id uint) (*models.Employee, error)
	Create(ctx context.Context, employee *models.Employee) error
	Update(ctx context.Context, employee *models.Employee) error
	Delete(ctx context.Context, id uint) error
	SearchByName(ctx context.Context, name string) ([]*models.Employee, error)
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{Q: query.Use(db)}
}

type employeeRepository struct {
	Q *query.Query
}

// SearchByName implements [EmployeeRepository].
func (r *employeeRepository) SearchByName(ctx context.Context, name string) ([]*models.Employee, error) {
	q := r.Q.Employee
	return q.WithContext(ctx).Where(
		q.FirstName.Like("%"+name+"%"),
	).Or(
		q.LastName.Like("%"+name+"%"),
	).Find()
}

// Create implements [EmployeeRepository].
func (r *employeeRepository) Create(ctx context.Context, employee *models.Employee) error {
	return r.Q.Employee.WithContext(ctx).Create(employee)
}

// Delete implements [EmployeeRepository].
func (r *employeeRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.Q.Employee.WithContext(ctx).Where(r.Q.Employee.EmployeeID.Eq(int32(id))).Delete(&models.Employee{})
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements [EmployeeRepository].
func (r *employeeRepository) FindAll(ctx context.Context) ([]*models.Employee, error) {
	var employees []*models.Employee
	employees, err := r.Q.Employee.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	return employees, err
}

// FindByID implements [EmployeeRepository].
func (r *employeeRepository) FindByID(ctx context.Context, id uint) (*models.Employee, error) {
	employee, err := r.Q.Employee.WithContext(ctx).Where(r.Q.Employee.EmployeeID.Eq(int32(id))).First()
	if err != nil {
		return nil, err
	}
	return employee, nil
}

// Update implements [EmployeeRepository].
func (r *employeeRepository) Update(ctx context.Context, employee *models.Employee) error {
	return r.Q.Employee.WithContext(ctx).Save(employee)
}