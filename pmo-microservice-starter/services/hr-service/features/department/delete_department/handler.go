package deletedapartment

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

type DeleteDepartmentHandler struct {
	repo      repository.DepartmentRepository
	validator *validators.DepartmentValidator
}

func NewDeleteDepartmentHandler(repo repository.DepartmentRepository, v *validators.DepartmentValidator) *DeleteDepartmentHandler {
	return &DeleteDepartmentHandler{
		repo:      repo,
		validator: v,
	}
}
func (h *DeleteDepartmentHandler) Handle(c *gin.Context) {

	// Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse[any]("Invalid department ID"))
		return
	}

	// Create request DTO
	req := &dto.DeleteDepartmentRequest{
		DepartmentID: int32(id),
	}

	// Validate
	if errs := h.validator.ValidateDeleteDepartment(req); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, response.ValidationError[any](errs))
		return
	}

	// Check if department exists
	dept, err := h.repo.GetByID(context.Background(), req.DepartmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse[any]("Failed to get department"))
		return
	}
	if dept == nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse[any]("Department not found"))
		return
	}

	// Delete
	if err := h.repo.Delete(context.Background(), req.DepartmentID); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse[any]("Failed to delete department"))
		return
	}

	// Return response
	c.JSON(http.StatusOK, response.SuccessResponse[any](nil, "Department deleted successfully"))
}
