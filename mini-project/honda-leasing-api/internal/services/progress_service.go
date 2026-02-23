package services

import (
	"context"
	"time"

	"github.com/codeid/honda-leasing-api/internal/dto"
	errs "github.com/codeid/honda-leasing-api/internal/errors"
	"github.com/codeid/honda-leasing-api/internal/repositories"
)

type OrderProgressServiceInterface interface {
	GetOrderProgress(ctx context.Context, contractID int64) ([]dto.OrderProgressResponse, error)
	UpdateTaskStatus(ctx context.Context, taskID int64, status string) error
}

type orderProgressService struct {
	contractRepo repositories.LeasingContractRepository
	taskRepo     repositories.LeasingTaskRepository
}

func NewOrderProgressService(
	contractRepo repositories.LeasingContractRepository,
	taskRepo repositories.LeasingTaskRepository,
) OrderProgressServiceInterface {
	return &orderProgressService{
		contractRepo: contractRepo,
		taskRepo:     taskRepo,
	}
}

// GetOrderProgress implements [OrderProgressServiceInterface].
func (s *orderProgressService) GetOrderProgress(
	ctx context.Context,
	contractID int64,
) ([]dto.OrderProgressResponse, error) {

	contract, err := s.contractRepo.FindByID(ctx, contractID)
	if err != nil {
		return nil, err
	}
	if contract == nil {
		return nil, errs.ErrContractNotFound
	}

	tasks, err := s.taskRepo.FindByContractID(ctx, contractID)
	if err != nil {
		return nil, err
	}

	var responses []dto.OrderProgressResponse

	for _, t := range tasks {

		status := ""
		if t.Status != nil {
			status = *t.Status
		}

		isDone := status == "completed"

		var date *time.Time
		if t.ActualEnddate != nil {
			date = t.ActualEnddate
		} else if t.ActualStartdate != nil {
			date = t.ActualStartdate
		}

		responses = append(responses, dto.OrderProgressResponse{
			TaskID:   t.TaskID,
			TaskName: t.TaskName,
			Status:   status,
			Date:     date,
			IsDone:   isDone,
		})
	}

	return responses, nil
}

// UpdateTaskStatus implements [OrderProgressServiceInterface].
func (s *orderProgressService) UpdateTaskStatus(ctx context.Context, taskID int64, status string) error {

	task, err := s.taskRepo.FindByID(ctx, taskID)
	if err != nil {
		return err
	}
	if task == nil {
		return errs.ErrTaskNotFound
	}

	now := time.Now()

	task.Status = &status

	if status == "completed" {
		task.ActualEnddate = &now
	} else {
		task.ActualStartdate = &now
	}

	return s.taskRepo.Update(ctx, task)
}
