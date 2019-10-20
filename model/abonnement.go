package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Abonnement struct {
	gorm.Model
	Date_deb time.Time `json:"date_deb" binding:"required"`
	Date_fin time.Time `json:"date_fin" binding:"required"`
	Struc    Structure `gorm:"foreign" json:"stucture" binding:"required"`
}
