package dto

type DeliveryProgressResponse struct {
	ContractID   int64                  `json:"contract_id"`
	CustomerName string                 `json:"customer_name"`
	RequestDate  string                 `json:"request_date"`
	Status       string                 `json:"status"`
	MotorName    string                 `json:"motor_name"`
	HargaOTR     float64                `json:"harga_otr"`
	Evidences    []DeliveryEvidenceItem `json:"evidences,omitempty"`
}

type DeliveryEvidenceItem struct {
	EvidenceID int64  `json:"evidence_id"`
	FileURL    string `json:"file_url"`
	UploadedAt string `json:"uploaded_at"`
}

type UploadDeliveryEvidenceRequest struct {
	DeliveryID int64  `json:"delivery_id" binding:"required"`
	FileURL    string `json:"file_url" binding:"required"`
	UploadedBy int64  `json:"uploaded_by"`
}
