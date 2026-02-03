package handlers

import (
	"net/http"
	"strconv"

	"github.com/codeid/hr-api/internal/dto"
	"github.com/codeid/hr-api/internal/response"
	"github.com/codeid/hr-api/internal/services"
	"github.com/gin-gonic/gin"
)

// EmployeeHandler handles HTTP requests for employee
type EmployeeHandler struct {
	service services.EmployeeServiceInterface
}

// EmployeeHandler creates a new handler instance
func NewEmployeeHandler(svc services.EmployeeServiceInterface) *EmployeeHandler {
	return &EmployeeHandler{
		service: svc,
	}
}

// CreateEmployee handles POST /departments
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to create employe: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusCreated, "Employee created successfully", resp)
}

// GetEmployeeByID handles GET /departments/:id
func (h *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.FindByID(c.Request.Context(), uint(id))
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Employee not found: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "Get Employee successfully", resp)
}

// GetAllEmployee handles GET /departments
func (h *EmployeeHandler) GetAllEmployee(c *gin.Context) {
	resp, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch employee: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "", resp)
}

// UpdateDepartment handles PUT /departments/:id
func (h *EmployeeHandler) UpdateHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch employees: "+err.Error())
		return
	}

	var req dto.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch employees: "+err.Error())
		return
	}

	resp, err := h.service.Update(c.Request.Context(), uint(id), &req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch employee: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "", resp)
}

// DeleteDepartment handles DELETE /departments/:id
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.SendError(c, http.StatusNotFound, "Data not found:"+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "Employee deleted successfully", nil)
}

// SearchDepartments handles GET /departments/search?q=...
func (h *EmployeeHandler) SearchEmployees(c *gin.Context) {
	name := c.Query("q")
	if name == "" {
		response.SendError(c, http.StatusBadRequest, "search query required: ")
		return
	}

	resp, err := h.service.SearchByName(c.Request.Context(), name)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch employees: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "", resp)
}