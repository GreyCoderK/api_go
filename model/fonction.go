package model

import (
	"github.com/jinzhu/gorm"
)

type Fonction struct {
	gorm.Model
	Libelle string `gorm:"type:varchar(100);not null;unique" json:"libelle" binding:"required"`
}

type Fonctions []Fonction
