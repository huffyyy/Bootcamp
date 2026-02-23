package services

import (
	"context"
	"errors"

	"github.com/codeid/honda-leasing-api/internal/dto"
	"github.com/codeid/honda-leasing-api/internal/repositories"
)

type SimulationServiceInterface interface {
	Simulate(ctx context.Context, req *dto.SimulationRequest) (*dto.SimulationResponse, error)
}

type simulationService struct {
	motorRepo   repositories.MotorRepository
	productRepo repositories.LeasingProductRepository
}

func NewSimulationService(
	motorRepo repositories.MotorRepository,
	productRepo repositories.LeasingProductRepository,
) SimulationServiceInterface {
	return &simulationService{
		motorRepo:   motorRepo,
		productRepo: productRepo,
	}
}

func safeFloat64(val *float64) float64 {
	if val == nil {
		return 0
	}
	return *val
}

func (s *simulationService) Simulate(ctx context.Context, req *dto.SimulationRequest) (*dto.SimulationResponse, error) {

	motor, err := s.motorRepo.FindByID(ctx, int64(req.MotorID))
	if err != nil {
		return nil, err
	}
	if motor == nil {
		return nil, errors.New("motor not found")
	}
	if motor.HargaOtr <= 0 {
		return nil, errors.New("invalid harga OTR")
	}

	product, err := s.productRepo.FindByID(ctx, int64(req.ProductID))
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product not found")
	}

	dp := req.DP
	if req.AllDP {
		dp = motor.HargaOtr
	}

	if dp < 0 {
		return nil, errors.New("dp cannot be negative")
	}

	if dp > motor.HargaOtr {
		return nil, errors.New("dp cannot be greater than harga OTR")
	}

	if dp != motor.HargaOtr {
		dpPercent := (dp / motor.HargaOtr) * 100

		if dpPercent < product.DpPersenMin || dpPercent > product.DpPersenMax {
			return nil, errors.New("dp percentage out of allowed range")
		}
	}

	if dp == motor.HargaOtr {
		return &dto.SimulationResponse{
			MotorID:  req.MotorID,
			HargaOTR: motor.HargaOtr,
			DP:       dp,
			Options:  []dto.TenorSimulation{},
		}, nil
	}

	pokok := motor.HargaOtr - dp

	tenorInt := int(product.TenorBulan)
	if tenorInt <= 0 {
		return nil, errors.New("invalid tenor")
	}

	bungaRate := product.BungaFlat / 100
	totalBunga := pokok * bungaRate * float64(tenorInt)

	adminFee := safeFloat64(product.AdminFee)

	total := pokok + totalBunga + adminFee

	if total <= 0 {
		return nil, errors.New("invalid total calculation")
	}

	angsuran := total / float64(tenorInt)

	options := []dto.TenorSimulation{
		{
			Tenor:    tenorInt,
			Angsuran: angsuran,
		},
	}

	return &dto.SimulationResponse{
		MotorID:  req.MotorID,
		HargaOTR: motor.HargaOtr,
		DP:       dp,
		Options:  options,
	}, nil
}
