package repositories

import (
	"context"

	"github.com/codeid/honda-leasing-api/internal/domain/models"
	"github.com/codeid/honda-leasing-api/internal/domain/query"
	"gorm.io/gorm"
)

type LeasingProductRepository interface {
	FindAll(ctx context.Context) ([]*models.LeasingProduct, error)
	FindByID(ctx context.Context, id int64) (*models.LeasingProduct, error)
}

func NewLeasingProductRepository(db *gorm.DB) LeasingProductRepository {
	return &leasingProductRepository{Q: query.Use(db)}
}

type leasingProductRepository struct {
	Q *query.Query
}

// FindAll implements [LeasingProductRepository].
func (r *leasingProductRepository) FindAll(ctx context.Context) ([]*models.LeasingProduct, error) {
	return r.Q.LeasingProduct.WithContext(ctx).Find()
}

// FindByID implements [LeasingProductRepository].
func (r *leasingProductRepository) FindByID(ctx context.Context, id int64) (*models.LeasingProduct, error) {
	leasingProduct, err := r.Q.LeasingProduct.
		WithContext(ctx).
		Where(r.Q.LeasingProduct.ProductID.Eq(id)).
		First()

	if err != nil {
		return nil, err
	}
	return leasingProduct, nil
}
