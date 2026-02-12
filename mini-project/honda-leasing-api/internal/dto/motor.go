package dto

// MotorListResponse for output (hide internal fields if needed)
type MotorListResponse struct {
	MotorID   int64   `json:"motor_id"`
	NamaModel string  `json:"nama_model"`
	Merk      string  `json:"merk"`
	HargaOtr  float64 `json:"harga_otr"`
	MotyName  string  `json:"moty_name"`
	FileName  string  `json:"file_name"`
}

// MotorDetailResponse for output (hide internal fields if needed)
type MotorDetailResponse struct {
	MotorID   int64   `json:"motor_id"`
	NamaModel string  `json:"nama_model"`
	Merk      string  `json:"merk"`
	HargaOtr  float64 `json:"harga_otr"`
	FileName  string  `json:"file_name"`
}
