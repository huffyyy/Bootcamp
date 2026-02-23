package repositories

import (
	"context"

	"github.com/codeid/honda-leasing-api/internal/domain/models"
	"github.com/codeid/honda-leasing-api/internal/domain/query"
	"gorm.io/gorm"
)

type LeasingTaskAtrRepository interface {
	FindByTaskID(ctx context.Context, taskID int64) ([]*models.LeasingTaskAtr, error)
}

func NewLeasingTaskAtrRepository(db *gorm.DB) LeasingTaskAtrRepository {
	return &leasingTaskAtrRepository{Q: query.Use(db)}
}

type leasingTaskAtrRepository struct {
	Q *query.Query
}

// FindByTaskID implements [LeasingTaskAtrRepository].
func (r *leasingTaskAtrRepository) FindByTaskID(ctx context.Context, taskID int64) ([]*models.LeasingTaskAtr, error) {
	tasks, err := r.Q.LeasingTaskAtr.
		WithContext(ctx).
		Where(r.Q.LeasingTaskAtr.TasaID.Eq(taskID)).
		Find()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
