package model

import (
	"github.com/jinzhu/gorm"
)

type Structure struct {
	gorm.Model
	RaisonSocial string `gorm:"type:varchar(255);not null;unique" json:"raisonSocial" binding:"required"`
}

type Structures []Structure
