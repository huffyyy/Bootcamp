package dto

type SimulationRequest struct {
	MotorID uint    `json:"motor_id" validate:"required"`
	DP      float64 `json:"dp" validate:"required"`
	Tenor   int     `json:"tenor" validate:"required"`
}

type SimulationResponse struct {
	MonthlyInstallment float64 `json:"monthly_installment"`
	TotalPayment       float64 `json:"total_payment"`
}
