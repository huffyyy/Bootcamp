package repositories

import (
	"context"

	"github.com/codeid/honda-leasing-api/internal/domain/models"
	"github.com/codeid/honda-leasing-api/internal/domain/query"
	"gorm.io/gorm"
)

type MotorRepository interface {
	FindAll(ctx context.Context) ([]*models.Motor, error)
	FindByID(ctx context.Context, id int64) (*models.Motor, error)
	SearchByName(ctx context.Context, name string) ([]*models.Motor, error)
	FindByCategory(ctx context.Context, category string) ([]*models.Motor, error)
}

func NewMotorRepository(db *gorm.DB) MotorRepository {
	return &motorRepository{Q: query.Use(db)}
}

type motorRepository struct {
	Q *query.Query
}

// FindByCategory implements MotorRepository
func (r *motorRepository) FindByCategory(ctx context.Context, category string) ([]*models.Motor, error) {
	m := r.Q.Motor
	mt := r.Q.MotorType
	motors, err := m.
		WithContext(ctx).
		Join(mt, m.MotorMotyID.EqCol(mt.MotyID)).
		Where(mt.MotyName.Eq(category)).
		Find()
	if err != nil {
		return nil, err
	}
	return motors, nil
}

// FindAll implements [MotorRepository].
func (r *motorRepository) FindAll(ctx context.Context) ([]*models.Motor, error) {
	return r.Q.Motor.WithContext(ctx).Find()
}

// FindByID implements [MotorRepository].
func (r *motorRepository) FindByID(ctx context.Context, id int64) (*models.Motor, error) {
	motors, err := r.Q.Motor.
		WithContext(ctx).
		Where(r.Q.Motor.MotorID.Eq(id)).
		First()

	if err != nil {
		return nil, err
	}
	return motors, nil
}

// SearchByName implements [MotorRepository].
func (r *motorRepository) SearchByName(ctx context.Context, name string) ([]*models.Motor, error) {
	m := r.Q.Motor
	return m.WithContext(ctx).
		Where(m.NamaModel.Like("%" + name + "%")).
		Find()
}
