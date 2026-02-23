package handlers

import (
	"net/http"

	"github.com/codeid/honda-leasing-api/internal/dto"
	"github.com/codeid/honda-leasing-api/internal/response"
	"github.com/codeid/honda-leasing-api/internal/services"
	"github.com/gin-gonic/gin"
)

type SimulationHandler struct {
	service services.SimulationServiceInterface
}

func NewSimulationHandler(svc services.SimulationServiceInterface) *SimulationHandler {
	return &SimulationHandler{
		service: svc,
	}
}

// Simulate handles POST/simulations
func (h *SimulationHandler) Simulate(c *gin.Context) {
	var req dto.SimulationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "invalid request body")
		return
	}
	resp, err := h.service.Simulate(c.Request.Context(), &req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "Simulation success", resp)
}
