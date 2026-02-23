package repositories

import (
	"context"

	"github.com/codeid/honda-leasing-api/internal/domain/models"
	"github.com/codeid/honda-leasing-api/internal/domain/query"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]*models.Customer, error)
	FindByID(ctx context.Context, id int64) (*models.Customer, error)
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{Q: query.Use(db)}
}

type customerRepository struct {
	Q *query.Query
}

func (r *customerRepository) FindAll(ctx context.Context) ([]*models.Customer, error) {
	return r.Q.Customer.WithContext(ctx).Find()
}

func (r *customerRepository) FindByID(ctx context.Context, id int64) (*models.Customer, error) {
	customer, err := r.Q.Customer.
		WithContext(ctx).
		Where(r.Q.Customer.CustomerID.Eq(id)).
		First()

	if err != nil {
		return nil, err
	}
	return customer, nil
}
