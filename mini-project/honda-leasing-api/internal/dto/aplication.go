package dto

type CreateApplicationRequest struct {
	MotorID uint    `json:"motor_id" validate:"required"`
	DP      float64 `json:"dp" validate:"required"`
	Tenor   int     `json:"tenor" validate:"required"`
	PromoID *uint   `json:"promo_id omitempty"`
}

type CreateApplicationResponse struct {
	ApplicationID uint   `json:"application_id"`
	Status        string `json:"status"`
}

type ApplicationDetailResponse struct {
	ID       uint           `json:"id"`
	Status   string         `json:"status"`
	Progress []ProgressItem `json:"progress"`
}

type ProgressItem struct {
	Step string `json:"step"`
	Done bool   `json:"done"`
}
