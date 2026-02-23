package handlers

import (
	"net/http"
	"strconv"

	"github.com/codeid/honda-leasing-api/internal/dto"
	"github.com/codeid/honda-leasing-api/internal/response"
	"github.com/codeid/honda-leasing-api/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderProgressHandler struct {
	service services.OrderProgressServiceInterface
}

func NewOrderProgressHandler(svc services.OrderProgressServiceInterface) *OrderProgressHandler {
	return &OrderProgressHandler{
		service: svc,
	}
}

// GetOrderProgress handles GET /orders/:id
func (h *OrderProgressHandler) GetOrderProgress(c *gin.Context) {

	idStr := c.Param("id")
	contractID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.GetOrderProgress(c.Request.Context(), contractID)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch order progress: "+err.Error())
		return
	}

	response.SendResponse(c, http.StatusOK, "Get order progress successfully", resp)
}

// UpdateTaskStatus handles PATCH /tasks/:task_id
func (h *OrderProgressHandler) UpdateTaskStatus(c *gin.Context) {

	taskIDStr := c.Param("task_id")
	taskID, err := strconv.ParseInt(taskIDStr, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	var req dto.UpdateTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	err = h.service.UpdateTaskStatus(c.Request.Context(), taskID, req.Status)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to update task: "+err.Error())
		return
	}

	response.SendResponse(c, http.StatusOK, "Update task successfully", nil)
}
