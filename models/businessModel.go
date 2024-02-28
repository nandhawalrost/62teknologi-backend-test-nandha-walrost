package models

import (
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	Location                 string
	Latitude                 uint
	Longitude                uint
	Term                     string
	Radius                   uint
	Categories               string
	Locale                   string
	Price                    uint
	Open_now                 bool
	Open_at                  uint
	Attributes               string
	Sort_by                  string
	Device_platform          string
	Reservation_date         string
	Reservation_time         string
	Reservation_covers       uint
	Matches_party_size_param bool `gorm:"default:false"`
	Limit                    uint
	Offset                   uint
}
