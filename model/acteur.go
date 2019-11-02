package model

import (
	"github.com/jinzhu/gorm"
)

type Acteur struct {
	gorm.Model
	Nom      string    `gorm:"type:varchar(100);not null" json:"nom" binding:"required"`
	Prenom   string    `gorm:"type:varchar(100);not null" json:"prenom" binding:"required"`
	Statut   string    `gorm:"type:varchar(100);not null" json:"statut" binding:"required"`
	Email    string    `gorm:"type:varchar(100);not null;unique" json:"email" binding:"required"`
	Pwd      string    `gorm:"type:varchar(255);not null" json:"pwd"`
	Fonction Fonction  `json:"fonction" binding:"required"`
	Struc    Structure `json:"structure" binding:"required"`
}

type Acteurs []Acteur
