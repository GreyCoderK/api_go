package model

import (
	"github.com/jinzhu/gorm"
)

type BonLivraison struct {
	gorm.Model
	PrixAchat *uint     `gorm:"default:0" json:"prix_achat"`
	Quantite  *uint     `gorm:"default:0" json:"quantite_acheter"`
	Article   Article   `gorm:"foreignkey:artRefer"`
	Livraison Livraison `gorm:"foreignkey:livRefer"`
	Art       *uint     `json:"article" binding:"required"`
	Liv       *uint     `json:"livraison" binding:"required"`
}

type BonLivraisons []BonLivraison
