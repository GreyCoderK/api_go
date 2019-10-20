package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Facture struct {
	gorm.Model
	Date    time.Time `json:"date"`
	Montant *uint     `gorm:"default:0" json:"montant"`
	Act     Acteur    `json:"acteur" binding:"required"`
}
