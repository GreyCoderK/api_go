package model

import (
	"github.com/jinzhu/gorm"
)

type Position struct {
	gorm.Model
	Lat   float64   `gorm:"type:decimal(10,8)"`
	Lon   float64   `gorm:"type:decimal(10,8)"`
	Struc Structure `json:"structure" binding:"required"`
}
