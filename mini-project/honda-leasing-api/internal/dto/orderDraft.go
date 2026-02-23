package dto

type OrderDraftRequest struct {
	CustomerID uint    `json:"customer_id" validate:"required"`
	MotorID    uint    `json:"motor_id" validate:"required"`
	ProductID  uint    `json:"product_id" validate:"required"`
	DP         float64 `json:"dp"`
	AllDP      bool    `json:"all_dp"`
}

type OrderDraftResponse struct {
	ContractID      uint    `json:"contract_id"`
	ContractNumber  string  `json:"contract_number"`
	Status          string  `json:"status"`
	TotalPinjaman   float64 `json:"total_pinjaman"`
	CicilanPerBulan float64 `json:"cicilan_per_bulan"`
}
