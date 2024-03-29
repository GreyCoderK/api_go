package model

import (
	"github.com/jinzhu/gorm"
)

type Position struct {
	gorm.Model
	Lat        float64   `json:"lon" gorm:"type:decimal(10,8)"`
	Lon        float64   `json:"lat" gorm:"type:decimal(10,8)"`
	Structure  Structure `gorm:"foreignkey:StrucRefer"`
	StrucRefer uint      `json:"structure" binding:"required"`
}

type Positions []Position
