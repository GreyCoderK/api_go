package repositories

import (
	. "../model"
	"github.com/jinzhu/gorm"
)

type BonCommandeRepository struct {
	db *gorm.DB
}

func NewBonCommandeRepository(db *gorm.DB) *BonCommandeRepository {
	return &BonCommandeRepository{db: db}
}

func (r *BonCommandeRepository) Save(m *Fonction) RepositoryResult {
	err := r.db.Save(m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: m}
}
