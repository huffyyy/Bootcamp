package models

import "time"

type BaseModel struct {
	CreatedDate  time.Time `gorm:"column:created_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_date"`
	ModifiedDate time.Time `gorm:"column:modified_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"modified_date"`
}

type Region struct {
	RegionID   uint   `gorm:"primaryKey;column:region_id"  json:"region_id"`
	RegionName string `gorm:"column:region_name;type:varchar(25)" json:"region_name"`
	BaseModel         //embendded
}

func (Region) TableName() string { return "hr.regions" }