package dto

type MotorListResponse struct {
	ID       int32   `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}

type MotorDetailResponse struct {
	ID          int32    `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Features    []string `json:"features"`
	Description string   `json:"description"`
}
