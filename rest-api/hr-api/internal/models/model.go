package models

import "time"

// BaseModel Struct
type BaseModel struct {
	CreatedDate  time.Time `gorm:"column:created_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_date"`
	ModifiedDate time.Time `gorm:"column:modified_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"modified_date"`
}

type Region struct {
	RegionID   uint   `gorm:"column:region_id;primaryKey" json:"region_id"`
	RegionName string `gorm:"column:region_name" json:"region_name"`
}

func (Region) TableName() string {
	return "regions"
}

// Country Struct

type Country struct {
	CountryID   string `gorm:"column:country_id;primaryKey" json:"country_id"`
	CountryName string `gorm:"column:country_name" json:"country_name"`
	RegionID    uint   `gorm:"column:region_id" json:"region_id"`

	Region Region `gorm:"foreignKey:RegionID;references:RegionID" json:"region"`
}

func (Country) TableName() string {
	return "countries"
}

