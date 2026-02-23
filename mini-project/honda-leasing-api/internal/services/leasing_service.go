package services

import (
	"context"

	"github.com/codeid/honda-leasing-api/internal/dto"
	errs "github.com/codeid/honda-leasing-api/internal/errors"
	"github.com/codeid/honda-leasing-api/internal/repositories"
)

type LeasingServiceInterface interface {
	GetInboxByStatus(ctx context.Context, status string) ([]dto.LeasingInboxResponse, error)
	GetRequestDetail(ctx context.Context, contractID int64) (*dto.LeasingDetailResponse, error)
	GetProgressDetail(ctx context.Context, contractID int64) (*dto.LeasingProgressDetailResponse, error)
}

type leasingService struct {
	contractRepo repositories.LeasingContractRepository
	customerRepo repositories.CustomerRepository
	motorRepo    repositories.MotorRepository
	productRepo  repositories.LeasingProductRepository
	taskRepo     repositories.LeasingTaskRepository
	taskAtrRepo  repositories.LeasingTaskAtrRepository
}

func NewLeasingService(
	contractRepo repositories.LeasingContractRepository,
	customerRepo repositories.CustomerRepository,
	motorRepo repositories.MotorRepository,
	productRepo repositories.LeasingProductRepository,
	taskRepo repositories.LeasingTaskRepository,
	taskAtrRepo repositories.LeasingTaskAtrRepository,
) LeasingServiceInterface {
	return &leasingService{
		contractRepo: contractRepo,
		customerRepo: customerRepo,
		motorRepo:    motorRepo,
		productRepo:  productRepo,
		taskRepo:     taskRepo,
		taskAtrRepo:  taskAtrRepo,
	}
}

// GetRequestDetail implements [LeasingServiceInterface].
func (s *leasingService) GetRequestDetail(ctx context.Context, contractID int64) (*dto.LeasingDetailResponse, error) {

	contract, err := s.contractRepo.FindByID(ctx, contractID)
	if err != nil {
		return nil, err
	}
	if contract == nil {
		return nil, errs.ErrContractNotFound
	}

	customer, err := s.customerRepo.FindByID(ctx, contract.CustomerID)
	if err != nil {
		return nil, err
	}

	motor, err := s.motorRepo.FindByID(ctx, contract.MotorID)
	if err != nil {
		return nil, err
	}

	product, err := s.productRepo.FindByID(ctx, contract.ProductID)
	if err != nil {
		return nil, err
	}

	adminFee := 0.0
	if product.AdminFee != nil {
		adminFee = *product.AdminFee
	}

	fidusia := 200000.0
	materai := 10000.0
	insurance := 250000.0

	subTotal := contract.NilaiKendaraan -
		contract.DpDibayar +
		adminFee +
		fidusia +
		materai +
		insurance

	return &dto.LeasingDetailResponse{
		ContractID:   contract.ContractID,
		CustomerName: customer.NamaLengkap,
		RequestDate:  contract.RequestDate.Format("02-Jan-2006"),
		Status:       contract.Status,

		MotorName:   *motor.NamaModel,
		HargaOTR:    contract.NilaiKendaraan,
		DownPayment: contract.DpDibayar,
		AdminFee:    adminFee,
		Insurance:   insurance,
		Fidusia:     fidusia,
		Materai:     materai,
		SubTotal:    subTotal,
	}, nil
}

// GetInboxByStatus implements [leasingServiceInterface].
func (s *leasingService) GetInboxByStatus(ctx context.Context, status string) ([]dto.LeasingInboxResponse, error) {

	contracts, err := s.contractRepo.FindByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var responses []dto.LeasingInboxResponse

	for _, c := range contracts {

		customer, err := s.customerRepo.FindByID(ctx, c.CustomerID)
		if err != nil {
			return nil, err
		}

		motor, err := s.motorRepo.FindByID(ctx, c.MotorID)
		if err != nil {
			return nil, err
		}

		responses = append(responses, dto.LeasingInboxResponse{
			ContractID:   c.ContractID,
			CustomerName: customer.NamaLengkap,
			RequestDate:  c.RequestDate.Format("02-Jan-2006"),
			MotorName:    *motor.NamaModel,
			Status:       c.Status,
		})
	}

	return responses, nil
}

// GetProgressDetail implements [leasingServiceInterface].

func (s *leasingService) GetProgressDetail(ctx context.Context, contractID int64) (*dto.LeasingProgressDetailResponse, error) {

	contract, err := s.contractRepo.FindByID(ctx, contractID)
	if err != nil {
		return nil, err
	}
	if contract == nil {
		return nil, errs.ErrContractNotFound
	}

	// Ambil semua task berdasarkan contract
	tasks, err := s.taskRepo.FindByContractID(ctx, contractID)
	if err != nil {
		return nil, err
	}

	var taskResponses []dto.LeasingTaskProgressItem

	for _, t := range tasks {

		status := ""
		if t.Status != nil {
			status = *t.Status
		}

		taskItem := dto.LeasingTaskProgressItem{
			TaskID:   t.TaskID,
			TaskName: t.TaskName,
			Status:   status,
		}

		if t.ActualStartdate != nil {
			taskItem.StartDate = t.ActualStartdate.Format("02-Jan-2006")
		}
		if t.ActualEnddate != nil {
			taskItem.EndDate = t.ActualEnddate.Format("02-Jan-2006")
		}

		// Ambil subtask
		subtasks, err := s.taskAtrRepo.FindByTaskID(ctx, t.TaskID)
		if err != nil {
			return nil, err
		}

		for _, st := range subtasks {

			subStatus := ""
			if st.TasaStatus != nil {
				subStatus = *st.TasaStatus
			}

			taskItem.Attributes = append(taskItem.Attributes,
				dto.LeasingTaskSubItem{
					AttributeID: st.TasaID,
					Name:        st.TasaName,
					Status:      subStatus,
				},
			)
		}

		taskResponses = append(taskResponses, taskItem)
	}

	return &dto.LeasingProgressDetailResponse{
		ContractID: contract.ContractID,
		Status:     contract.Status,
		Tasks:      taskResponses,
	}, nil
}
