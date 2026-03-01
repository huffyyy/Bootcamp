package getdepartmentbyname

import (
	"context"
	"net/http"
	"pmo/internal/pkg/response"
	"pmo/services/hr-service/features/department/shared/dto"
	"pmo/services/hr-service/features/department/shared/repository"
	"pmo/services/hr-service/features/department/shared/validators"

	"github.com/gin-gonic/gin"
)

type GetDepartmentByNameHandler struct {
	repo      repository.DepartmentRepository
	validator *validators.DepartmentValidator
}

func NewGetDepartmentByNameHandler(repo repository.DepartmentRepository, v *validators.DepartmentValidator) *GetDepartmentByNameHandler {
	return &GetDepartmentByNameHandler{
		repo:      repo,
		validator: v,
	}
}
func (h *GetDepartmentByNameHandler) Handle(c *gin.Context) {

	// Get query parameter
	var req dto.GetDepartmentByNameRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse[any]("Invalid query parameters"))
		return
	}

	// Validate
	if errs := h.validator.ValidateGetDepartmentByName(&req); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, response.ValidationError[any](errs))
		return
	}

	// Execute query
	departments, err := h.repo.GetByName(context.Background(), req.DepartmentName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse[any]("Failed to get departments"))
		return
	}

	// Return response
	resp := dto.ToDepartmentResponses(departments)
	if len(resp) == 0 {
		c.JSON(http.StatusOK, response.SuccessResponse([]dto.DepartmentResponse{}, "No departments found"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp, "Departments retrieved successfully"))
}
