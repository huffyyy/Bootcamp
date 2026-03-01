package getalldepartment

import (
	"context"
	"math"
	"net/http"
	"pmo/internal/pkg/response"
	"pmo/services/hr-service/features/department/shared/dto"
	"pmo/services/hr-service/features/department/shared/repository"
	"pmo/services/hr-service/features/department/shared/validators"

	"github.com/gin-gonic/gin"
)

type GetAllDepartmentsHandler struct {
	repo      repository.DepartmentRepository
	validator *validators.DepartmentValidator
}

func NewGetAllDepartmentsHandler(repo repository.DepartmentRepository, v *validators.DepartmentValidator) *GetAllDepartmentsHandler {
	return &GetAllDepartmentsHandler{
		repo:      repo,
		validator: v,
	}
}
func (h *GetAllDepartmentsHandler) Handle(c *gin.Context) {
	// Bind query parameters
	var req dto.GetAllDepartmentsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse[any]("Invalid query parameters"))
		return
	}
	// Validate
	if errs := h.validator.ValidateGetAllDepartments(&req); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, response.ValidationError[any](errs))
		return
	}
	// Create query
	query := NewGetAllDepartmentsQuery(
		req.Page,
		req.PageSize,
		req.SortBy,
		req.SortDir,
		req.DepartmentName,
		req.LocationID,
	)
	// Execute query
	departments, total, err := h.repo.FindAll(context.Background(),
		query.ToRepositoryParams())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse[any]("Failed to get departments: "+err.Error()))
		return
	}
	// Calculate pagination metadata
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))
	// Create filters for meta
	filters := []response.Filter{}
	if req.DepartmentName != "" {
		filters = append(filters, response.Filter{
			Field:    "department_name",
			Operator: "contains",
			Value:    req.DepartmentName,
		})
	}
	if req.LocationID != nil {
		filters = append(filters, response.Filter{
			Field:    "location_id",
			Operator: "eq",
			Value:    *req.LocationID,
		})
	}
	// Create meta
	meta := &response.Meta{
		CurrentPage: req.Page,
		PageSize:    req.PageSize,
		TotalPages:  totalPages,
		TotalItems:  int(total),
		SortBy:      req.SortBy,
		SortDir:     req.SortDir,
		Filters:     filters,
	}
	// Create response
	resp := dto.ToDepartmentListResponse(departments, total)
	// Return response with meta
	c.JSON(http.StatusOK, response.SuccessResponseWithMeta(resp, "Departments retrieved successfully", meta))
}
