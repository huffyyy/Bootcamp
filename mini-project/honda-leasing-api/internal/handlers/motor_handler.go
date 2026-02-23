package handlers

import (
	"net/http"
	"strconv"

	"github.com/codeid/honda-leasing-api/internal/response"
	"github.com/codeid/honda-leasing-api/internal/services"
	"github.com/gin-gonic/gin"
)

type MotorHandler struct {
	service services.MotorServiceInterface
}

func NewMotorHandler(svc services.MotorServiceInterface) *MotorHandler {
	return &MotorHandler{
		service: svc,
	}
}

// GetMotorByID handles GET /motors/:id
func (h *MotorHandler) GetMotorByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.FindByID(c.Request.Context(), int64(id))
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Motor not found: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "Get Motor successfully", resp)
}

// GetAllMotor handles GET /motors
func (h *MotorHandler) GetAllMotor(c *gin.Context) {
	resp, err := h.service.FindAll(c.Request.Context())
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch Motor: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "", resp)
}

// SearchMotors handles GET /motors/search?q=
func (h *MotorHandler) SearchMotors(c *gin.Context) {
	name := c.Query("q")
	if name == "" {
		response.SendError(c, http.StatusBadRequest, "search query required: ")
		return
	}

	resp, err := h.service.SearchByName(c.Request.Context(), name)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch motors: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "", resp)
}

// GetMotorByCategory handles GET /motors/category?category=
func (h *MotorHandler) GetMotorByCategory(c *gin.Context) {
	category := c.Query("category")
	if category == "" {
		response.SendError(c, http.StatusBadRequest, "category query required")
		return
	}

	resp, err := h.service.FindByCategory(c.Request.Context(), category)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch motors: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "Get motors by category successfully", resp)
}
