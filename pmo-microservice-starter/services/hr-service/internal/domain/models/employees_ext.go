package models

import "time"

type EmployeeExt struct {
	EmployeeID   int32           `gorm:"primaryKey;column:employee_id"`
	FirstName    *string         `gorm:"column:first_name"`
	LastName     string          `gorm:"column:last_name"`
	Email        string          `gorm:"column:email"`
	PhoneNumber  *string         `gorm:"column:phone_number"`
	HireDate     time.Time       `gorm:"column:hire_date"`
	JobID        int32           `gorm:"column:job_id"`
	Salary       float64         `gorm:"column:salary"`
	ManagerID    *int32          `gorm:"column:manager_id"`
	DepartmentID *int32          `gorm:"column:department_id"`
	Photos       []EmployeePhoto `gorm:"foreignKey:EmployeeID"`
}

type EmployeePhoto struct {
	EphoID     int32   `gorm:"primaryKey;column:epho_id"`
	EmployeeID int32   `gorm:"column:employee_id"`
	FileName   *string `gorm:"column:file_name"`
	FileSize   *int64  `gorm:"column:file_size"`
	FileType   *string `gorm:"column:file_type"`
	FileURL    *string `gorm:"column:file_url"`
	IsPrimary  *bool   `gorm:"column:is_primary"`
}

func (EmployeeExt) TableName() string {
	return "hr.employees"
}

func (EmployeePhoto) TableName() string {
	return "hr.employee_photos"
}
