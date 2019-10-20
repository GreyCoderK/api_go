package repositories

import (
	. "../model"
	"github.com/jinzhu/gorm"
)

type BonLivraisonRepository struct {
	db *gorm.DB
}

func NewBonLivraisonRepository(db *gorm.DB) *BonLivraisonRepository {
	return &BonLivraisonRepository{db: db}
}

func (r *BonLivraisonRepository) Save(m *Fonction) RepositoryResult {
	err := r.db.Save(m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: m}
}
