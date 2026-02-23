package repositories

import (
	"context"

	"github.com/codeid/honda-leasing-api/internal/domain/models"
	"github.com/codeid/honda-leasing-api/internal/domain/query"
	"gorm.io/gorm"
)

type LeasingContractRepository interface {
	FindAll(ctx context.Context) ([]*models.LeasingContract, error)
	FindByID(ctx context.Context, id int64) (*models.LeasingContract, error)
	Create(ctx context.Context, leasingContract *models.LeasingContract) error
	FindByStatus(ctx context.Context, status string) ([]*models.LeasingContract, error)
}

func NewLeasingContractRepository(db *gorm.DB) LeasingContractRepository {
	return &leasingContractRepository{Q: query.Use(db)}
}

type leasingContractRepository struct {
	Q *query.Query
}

// FindByStatus implements [LeasingContractRepository].
func (r *leasingContractRepository) FindByStatus(ctx context.Context, status string) ([]*models.LeasingContract, error) {

	contracts, err := r.Q.LeasingContract.
		WithContext(ctx).
		Where(r.Q.LeasingContract.Status.Eq(status)).
		Find()

	if err != nil {
		return nil, err
	}

	return contracts, nil
}

// Create implements [LeasingContractRepository].
func (r *leasingContractRepository) Create(ctx context.Context, leasingContract *models.LeasingContract) error {
	return r.Q.LeasingContract.WithContext(ctx).Create(leasingContract)
}

// FindAll implements [LeasingContractRepository].
func (r *leasingContractRepository) FindAll(ctx context.Context) ([]*models.LeasingContract, error) {
	return r.Q.LeasingContract.WithContext(ctx).Find()
}

// FindByID implements [LeasingContractRepository].
func (r *leasingContractRepository) FindByID(ctx context.Context, id int64) (*models.LeasingContract, error) {
	leasingContract, err := r.Q.LeasingContract.
		WithContext(ctx).
		Where(r.Q.LeasingContract.ContractID.Eq(id)).
		First()

	if err != nil {
		return nil, err
	}
	return leasingContract, nil
}
