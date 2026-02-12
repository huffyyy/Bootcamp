package services

import (
	"context"
	"strings"

	"github.com/codeid/honda-leasing-api/internal/dto"
	"github.com/codeid/honda-leasing-api/internal/repositories"
	"github.com/go-playground/validator/v10"
)

type MotorServiceInterface interface {
	FindByID(ctx context.Context, id int64) (*dto.MotorDetailResponse, error)
	GetAll(ctx context.Context) ([]dto.MotorListResponse, error)
	SearchByName(ctx context.Context, name string) ([]dto.MotorListResponse, error)
	FindByCategory(ctx context.Context, category string) ([]dto.MotorListResponse, error)
}

type motorService struct {
	repo     repositories.MotorRepository
	validate *validator.Validate
}

func NewMotorService(repo repositories.MotorRepository) MotorServiceInterface {
	return &motorService{
		repo:     repo,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

// helper to safely dereference optional fields
func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func shortModelName(full *string) string {
	if full == nil || *full == "" {
		return ""
	}

	parts := strings.Split(*full, " ")
	return parts[0]
}

// FindByID implements [MotorServiceInterface].
func (s *motorService) FindByID(ctx context.Context, id int64) (*dto.MotorDetailResponse, error) {
	mtr, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.MotorDetailResponse{
		MotorID:   mtr.MotorID,
		NamaModel: safeString(mtr.NamaModel),
		Merk:      mtr.Merk,
		HargaOtr:  mtr.HargaOtr,
		FileName:  "",
	}, nil

}

// GetAll implements [MotorServiceInterface].
func (s *motorService) GetAll(ctx context.Context) ([]dto.MotorListResponse, error) {
	mtrs, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var response []dto.MotorListResponse
	for _, mtr := range mtrs {
		response = append(response, dto.MotorListResponse{
			MotorID:   mtr.MotorID,
			NamaModel: shortModelName(mtr.NamaModel),
			Merk:      mtr.Merk,
			HargaOtr:  mtr.HargaOtr,
			MotyName:  safeString(mtr.MotorType),
			FileName:  "",
		})
	}
	return response, nil
}

// SearchByName implements [MotorServiceInterface].
func (s *motorService) SearchByName(ctx context.Context, name string) ([]dto.MotorListResponse, error) {
	if name == "" {
		return s.GetAll(ctx)
	}

	mtrs, err := s.repo.SearchByName(ctx, name)
	if err != nil {
		return nil, err
	}

	var responses []dto.MotorListResponse
	for _, mtr := range mtrs {
		responses = append(responses, dto.MotorListResponse{
			MotorID:   mtr.MotorID,
			NamaModel: shortModelName(mtr.NamaModel),
			Merk:      mtr.Merk,
			HargaOtr:  mtr.HargaOtr,
			MotyName:  safeString(mtr.MotorType),
			FileName:  "",
		})
	}
	return responses, nil
}

// FindByCategory implements [MotorServiceInterface].
func (s *motorService) FindByCategory(ctx context.Context, category string) ([]dto.MotorListResponse, error) {
	mtrs, err := s.repo.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	var responses []dto.MotorListResponse
	for _, mtr := range mtrs {
		responses = append(responses, dto.MotorListResponse{
			MotorID:   mtr.MotorID,
			NamaModel: shortModelName(mtr.NamaModel),
			Merk:      mtr.Merk,
			HargaOtr:  mtr.HargaOtr,
			MotyName:  safeString(mtr.MotorType),
			FileName:  "",
		})
	}
	return responses, nil
}
