package getdepartmentbyid

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

type GetDepartmentByIDHandler struct {
	repo      repository.DepartmentRepository
	validator *validators.DepartmentValidator
}

func NewGetDepartmentByIDHandler(repo repository.DepartmentRepository, v *validators.DepartmentValidator) *GetDepartmentByIDHandler {
	return &GetDepartmentByIDHandler{
		repo:      repo,
		validator: v,
	}
}
func (h *GetDepartmentByIDHandler) Handle(c *gin.Context) {

	// Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse[any]("Invalid department ID"))
		return
	}

	// Create request DTO
	req := &dto.GetDepartmentByIDRequest{
		DepartmentID: int32(id),
	}

	// Validate
	if errs := h.validator.ValidateGetDepartmentByID(req); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, response.ValidationError[any](errs))
		return
	}

	// Execute query
	dept, err := h.repo.GetByID(context.Background(), req.DepartmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse[any]("Failed to get department"))
		return
	}
	if dept == nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse[any]("Department not found"))
		return
	}

	// Return response
	resp := dto.ToDepartmentResponse(dept)
	c.JSON(http.StatusOK, response.SuccessResponse(resp, "Department retrieved successfully"))
}
