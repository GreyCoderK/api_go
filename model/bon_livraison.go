package model

import (
	"github.com/jinzhu/gorm"
)

type BonLivraison struct {
	gorm.Model
	PrixAchat *uint     `gorm:"default:0" json:"prix_achat"`
	Quantite  *uint     `gorm:"default:0" json:"quantite_acheter"`
	Art       []Article `json:"article"`
	Liv       Livraison `json:"livraison" binding:"required"`
}
