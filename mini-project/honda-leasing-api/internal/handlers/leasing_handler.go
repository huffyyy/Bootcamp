package handlers

import (
	"net/http"
	"strconv"

	"github.com/codeid/honda-leasing-api/internal/response"
	"github.com/codeid/honda-leasing-api/internal/services"
	"github.com/gin-gonic/gin"
)

type LeasingHandler struct {
	service services.LeasingServiceInterface
}

func NewLeasingHandler(svc services.LeasingServiceInterface) *LeasingHandler {
	return &LeasingHandler{
		service: svc,
	}
}

// GetInboxByStatus handles GET /leasing/inbox?status=request
func (h *LeasingHandler) GetInboxByStatus(c *gin.Context) {

	status := c.Query("status")
	if status == "" {
		response.SendError(c, http.StatusBadRequest, "status query required")
		return
	}

	resp, err := h.service.GetInboxByStatus(c.Request.Context(), status)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch inbox: "+err.Error())
		return
	}

	response.SendResponse(c, http.StatusOK, "Get inbox successfully", resp)
}

// GetRequestDetail handles GET /leasing/contracts/:id
func (h *LeasingHandler) GetRequestDetail(c *gin.Context) {

	idStr := c.Param("id")
	contractID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.GetRequestDetail(c.Request.Context(), contractID)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch request detail: "+err.Error())
		return
	}

	response.SendResponse(c, http.StatusOK, "Get request detail successfully", resp)
}

// GetProgressDetail handles GET /leasing/:id/progress-detail
func (h *LeasingHandler) GetProgressDetail(c *gin.Context) {

	idStr := c.Param("id")
	contractID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.GetProgressDetail(c.Request.Context(), contractID)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch progress detail: "+err.Error())
		return
	}

	response.SendResponse(c, http.StatusOK, "Get progress detail successfully", resp)
}
