package repositories

import (
	"context"

	"github.com/codeid/honda-leasing-api/internal/domain/models"
	"github.com/codeid/honda-leasing-api/internal/domain/query"
	"gorm.io/gorm"
)

type LeasingTaskRepository interface {
	FindByID(ctx context.Context, id int64) (*models.LeasingTask, error)
	FindByContractID(ctx context.Context, contractID int64) ([]*models.LeasingTask, error)
	Create(ctx context.Context, leasingTask *models.LeasingTask) error
	Update(ctx context.Context, leasingTask *models.LeasingTask) error
}

func NewLeasingTaskRepository(db *gorm.DB) LeasingTaskRepository {
	return &leasingTaskRepository{Q: query.Use(db)}
}

type leasingTaskRepository struct {
	Q *query.Query
}

// FindByContractID implements [LeasingTaskRepository].
func (r *leasingTaskRepository) FindByContractID(ctx context.Context, contractID int64) ([]*models.LeasingTask, error) {

	tasks, err := r.Q.LeasingTask.
		WithContext(ctx).
		Where(r.Q.LeasingTask.ContractID.Eq(contractID)).
		Order(r.Q.LeasingTask.SequenceNo.Asc()).
		Find()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Create implements [LeasingTaskRepository].
func (r *leasingTaskRepository) Create(ctx context.Context, leasingTasks *models.LeasingTask) error {
	return r.Q.LeasingTask.WithContext(ctx).Create(leasingTasks)
}

// FindByID implements [LeasingTaskRepository].
func (r *leasingTaskRepository) FindByID(ctx context.Context, id int64) (*models.LeasingTask, error) {
	leasingTask, err := r.Q.LeasingTask.
		WithContext(ctx).
		Where(r.Q.LeasingTask.TaskID.Eq(id)).
		First()

	if err != nil {
		return nil, err
	}
	return leasingTask, nil
}

// Update implements [LeasingTaskRepository].
func (r *leasingTaskRepository) Update(ctx context.Context, leasingTask *models.LeasingTask) error {
	return r.Q.LeasingTask.WithContext(ctx).Save(leasingTask)
}
