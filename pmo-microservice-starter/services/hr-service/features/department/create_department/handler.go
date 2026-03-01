package createdepartment

import (
	"context"
	"errors"
	"net/http"
	"pmo/internal/pkg/response"
	"pmo/services/hr-service/features/department/shared/dto"
	"pmo/services/hr-service/features/department/shared/repository"
	"pmo/services/hr-service/features/department/shared/validators"

	"github.com/gin-gonic/gin"
)

type CreateDepartmentHandler struct {
	repo      repository.DepartmentRepository
	validator *validators.DepartmentValidator
}

func NewCreateDepartmentHandler(repo repository.DepartmentRepository, v *validators.DepartmentValidator) *CreateDepartmentHandler {
	return &CreateDepartmentHandler{
		repo:      repo,
		validator: v,
	}
}
func (h *CreateDepartmentHandler) Handle(c *gin.Context) {
	var req dto.CreateDepartmentRequest
	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse[any]("Invalid request format"))
		return
	}
	// Validate
	if errs := h.validator.ValidateCreateDepartment(&req); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, response.ValidationError[any](errs))
		return
	}
	// Create command
	cmd := NewCreateDepartmentCommand(req.DepartmentName, req.LocationID)
	// Execute
	dept := cmd.ToModel()
	if err := h.repo.Create(context.Background(), dept); err != nil {
		// Check for duplicate name error
		if errors.Is(err, repository.ErrDuplicateName) {
			c.JSON(http.StatusConflict, response.ErrorResponse[any]("Department name already exists"))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorResponse[any]("Failed to create department"))
		return
	}
	// Return response
	resp := dto.ToDepartmentResponse(dept)
	c.JSON(http.StatusCreated, response.SuccessResponse(resp, "Department created successfully"))
}
