package handlers

import (
	"net/http"

	"github.com/codeid/honda-leasing-api/internal/dto"
	"github.com/codeid/honda-leasing-api/internal/response"
	"github.com/codeid/honda-leasing-api/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderDraftHandler struct {
	service services.OrderDraftServiceInterface
}

func NewOrderDraftHandler(svc services.OrderDraftServiceInterface) *OrderDraftHandler {
	return &OrderDraftHandler{
		service: svc,
	}
}

// CreateOrderDraft handles POST /orders
func (h *OrderDraftHandler) CreateOrderDraft(c *gin.Context) {
	var req dto.OrderDraftRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.CreateOrderDraft(c.Request.Context(), &req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to create draft: "+err.Error())
		return
	}

	response.SendResponse(c, http.StatusCreated, "Draft created successfully", resp)
}
