package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Livraison struct {
	gorm.Model
	Date    time.Time `json:"date" binding:"required"`
	Montant *uint     `gorm:"default:0" json:"montant"`
}

type Livraisons []Livraison
