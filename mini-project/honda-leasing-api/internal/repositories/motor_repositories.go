package repositories

import (
	"context"

	"github.com/codeid/honda-leasing-api/internal/domain/models"
	"github.com/codeid/honda-leasing-api/internal/model"
	"gorm.io/gorm"
)

type MotorRepository interface {
	FindAll(ctx context.Context) ([]*model.Motor, error)
	FindByID(ctx context.Context, id int64) (*model.Motor, error)
	SearchByName(ctx context.Context, name string) ([]*model.Motor, error)
	FindByCategory(ctx context.Context, category string) ([]*model.Motor, error)
}

func NewMotorRepository(db *gorm.DB) MotorRepository {
	return &motorRepository{Q: models.Use(db)}
}

type motorRepository struct {
	Q *models.Query
}

// FindByCategory implements MotorRepository
func (r *motorRepository) FindByCategory(ctx context.Context, category string) ([]*model.Motor, error) {
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
func (r *motorRepository) FindAll(ctx context.Context) ([]*model.Motor, error) {
	return r.Q.Motor.WithContext(ctx).Find()
}

// FindByID implements [MotorRepository].
func (r *motorRepository) FindByID(ctx context.Context, id int64) (*model.Motor, error) {
	return r.Q.Motor.
		WithContext(ctx).
		Where(r.Q.Motor.MotorID.Eq(id)).
		First()
}

// SearchByName implements [MotorRepository].
func (r *motorRepository) SearchByName(ctx context.Context, name string) ([]*model.Motor, error) {
	m := r.Q.Motor
	return m.WithContext(ctx).
		Where(m.NamaModel.Like("%" + name + "%")).
		Find()
}
