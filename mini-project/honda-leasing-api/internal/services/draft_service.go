package services

import (
	"context"
	"fmt"
	"time"

	"github.com/codeid/honda-leasing-api/internal/domain/models"
	"github.com/codeid/honda-leasing-api/internal/dto"
	errs "github.com/codeid/honda-leasing-api/internal/errors"
	"github.com/codeid/honda-leasing-api/internal/repositories"
)

type OrderDraftServiceInterface interface {
	CreateOrderDraft(ctx context.Context, req *dto.OrderDraftRequest) (*dto.OrderDraftResponse, error)
}

type orderDraftService struct {
	motorRepo    repositories.MotorRepository
	productRepo  repositories.LeasingProductRepository
	customerRepo repositories.CustomerRepository
	contractRepo repositories.LeasingContractRepository
}

func NewOrderDraftService(
	motorRepo repositories.MotorRepository,
	productRepo repositories.LeasingProductRepository,
	customerRepo repositories.CustomerRepository,
	contractRepo repositories.LeasingContractRepository,
) OrderDraftServiceInterface {
	return &orderDraftService{
		motorRepo:    motorRepo,
		productRepo:  productRepo,
		customerRepo: customerRepo,
		contractRepo: contractRepo,
	}
}

func (s *orderDraftService) CreateOrderDraft(ctx context.Context, req *dto.OrderDraftRequest) (*dto.OrderDraftResponse, error) {

	motor, err := s.motorRepo.FindByID(ctx, int64(req.MotorID))
	if err != nil {
		return nil, err
	}
	if motor == nil {
		return nil, errs.ErrMotorNotFound
	}

	product, err := s.productRepo.FindByID(ctx, int64(req.ProductID))
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errs.ErrProductNotFound
	}

	customer, err := s.customerRepo.FindByID(ctx, int64(req.CustomerID))
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errs.ErrCustomerNotFound
	}

	dp := req.DP
	if req.AllDP {
		dp = motor.HargaOtr
	}

	if dp < 0 || dp > motor.HargaOtr {
		return nil, errs.ErrInvalidDP
	}

	pokok := motor.HargaOtr - dp

	tenor := int(product.TenorBulan)
	if tenor <= 0 {
		return nil, errs.ErrInvalidTenor
	}

	bungaRate := product.BungaFlat / 100
	totalBunga := pokok * bungaRate * float64(tenor)

	adminFee := 0.0
	if product.AdminFee != nil {
		adminFee = *product.AdminFee
	}

	totalPinjaman := pokok + totalBunga + adminFee
	cicilanPerBulan := totalPinjaman / float64(tenor)

	contractNumber := fmt.Sprintf("LSG-%d-%d", time.Now().Year(), time.Now().Unix())

	now := time.Now()
	contract := &models.LeasingContract{
		ContractNumber:  &contractNumber,
		RequestDate:     time.Now(),
		TenorBulan:      int16(tenor),
		NilaiKendaraan:  motor.HargaOtr,
		DpDibayar:       dp,
		PokokPinjaman:   pokok,
		TotalPinjaman:   totalPinjaman,
		CicilanPerBulan: cicilanPerBulan,
		Status:          "draft",
		CustomerID:      customer.CustomerID,
		MotorID:         motor.MotorID,
		ProductID:       product.ProductID,
		CreatedAt:       &now,
	}

	if err := s.contractRepo.Create(ctx, contract); err != nil {
		return nil, errs.ErrCreateContract
	}

	return &dto.OrderDraftResponse{
		ContractID:      uint(contract.ContractID),
		ContractNumber:  contractNumber,
		Status:          "draft",
		TotalPinjaman:   totalPinjaman,
		CicilanPerBulan: cicilanPerBulan,
	}, nil
}
