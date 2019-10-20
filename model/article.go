package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Libelle      string           `gorm:"type:varchar(255);not null;unique" json:"libelle" binding:"required"`
	PrixUnitaire *uint            `gorm:"default:0" json:"prix_unitaire"`
	QteSeuil     *uint            `gorm:"default:0" json:"quantite_en_stock"`
	Categorie    CategorieArticle `json:"categorie" binding:"required"`
}
