package dto

type LeasingInboxResponse struct {
	ContractID   int64  `json:"contract_id"`
	CustomerName string `json:"customer_name"`
	RequestDate  string `json:"request_date"`
	MotorName    string `json:"motor_name"`
	MotorImage   string `json:"motor_image"`
	Status       string `json:"status"`
}

type LeasingDetailResponse struct {
	ContractID   int64   `json:"contract_id"`
	CustomerName string  `json:"customer_name"`
	RequestDate  string  `json:"request_date"`
	Status       string  `json:"status"`
	MotorName    string  `json:"motor_name"`
	HargaOTR     float64 `json:"harga_otr"`
	DownPayment  float64 `json:"down_payment"`
	AdminFee     float64 `json:"admin_fee"`
	Insurance    float64 `json:"insurance"`
	Fidusia      float64 `json:"fidusia"`
	Materai      float64 `json:"materai"`
	SubTotal     float64 `json:"sub_total"`
}

type LeasingProgressDetailResponse struct {
	ContractID int64                     `json:"contract_id"`
	Status     string                    `json:"status"`
	Tasks      []LeasingTaskProgressItem `json:"tasks"`
}

type LeasingTaskProgressItem struct {
	TaskID     int64                `json:"task_id"`
	TaskName   string               `json:"task_name"`
	Status     string               `json:"status"`
	StartDate  string               `json:"start_date,omitempty"`
	EndDate    string               `json:"end_date,omitempty"`
	Attributes []LeasingTaskSubItem `json:"attributes,omitempty"`
}

type LeasingTaskSubItem struct {
	AttributeID int64  `json:"attribute_id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
}
