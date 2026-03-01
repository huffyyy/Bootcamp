package updatedepartment

import (
	"context"
	"net/http"
	"pmo/internal/pkg/response"
	"pmo/services/hr-service/features/department/shared/dto"
	"pmo/services/hr-service/features/department/shared/repository"
	"pmo/services/hr-service/features/department/shared/validators"

	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateDepartmentHandler struct {
	repo      repository.DepartmentRepository
	validator *validators.DepartmentValidator
}

func NewUpdateDepartmentHandler(repo repository.DepartmentRepository, v *validators.DepartmentValidator) *UpdateDepartmentHandler {
	return &UpdateDepartmentHandler{
		repo:      repo,
		validator: v,
	}
}
func (h *UpdateDepartmentHandler) Handle(c *gin.Context) {

	// Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse[any]("Invalid department ID"))
		return
	}

	// Bind JSON
	var req dto.UpdateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse[any]("Invalid request format"))
		return
	}
	// Validate
	if errs := h.validator.ValidateUpdateDepartment(&req); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, response.ValidationError[any](errs))
		return
	}

	// Get existing department
	dept, err := h.repo.GetByID(context.Background(), int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse[any]("Failed to get department"))
		return
	}
	if dept == nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse[any]("Department not found"))
		return
	}

	// Create and apply command
	cmd := NewUpdateDepartmentCommand(int32(id), req.DepartmentName,
		req.LocationID)
	cmd.ApplyToModel(dept)

	// Update
	if err := h.repo.Update(context.Background(), dept); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse[any]("Failed to update department"))
		return
	}

	// Return response
	resp := dto.ToDepartmentResponse(dept)
	c.JSON(http.StatusOK, response.SuccessResponse(resp, "Department updated successfully"))
}
