package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Facture struct {
	gorm.Model
	Date     time.Time `json:"date"`
	Montant  *uint     `gorm:"default:0" json:"montant"`
	Acteur   Acteur    `gorm:"foreignkey:ActRefer"`
	ActRefer *uint     `json:"acteur" binding:"required"`
}

type Factures []Facture
