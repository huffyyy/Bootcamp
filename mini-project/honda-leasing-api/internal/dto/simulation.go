package dto

type SimulationRequest struct {
	MotorID   uint    `json:"motor_id" validate:"required"`
	ProductID uint    `json:"product_id" validate:"required"`
	DP        float64 `json:"dp"`
	AllDP     bool    `json:"all_dp"`
}

type TenorSimulation struct {
	Tenor    int     `json:"tenor"`
	Angsuran float64 `json:"angsuran"`
}

type SimulationResponse struct {
	MotorID  uint              `json:"motor_id"`
	HargaOTR float64           `json:"harga_otr"`
	DP       float64           `json:"dp"`
	Options  []TenorSimulation `json:"options"`
}
