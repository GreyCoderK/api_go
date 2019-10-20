package repositories

import (
	. "../model"
	"github.com/jinzhu/gorm"
)

type FonctionRepository struct {
	db *gorm.DB
}

func NewFonctionRepository(db *gorm.DB) *FonctionRepository {
	return &FonctionRepository{db: db}
}

func (r *FonctionRepository) Save(m *Fonction) RepositoryResult {
	err := r.db.Save(m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: m}
}
