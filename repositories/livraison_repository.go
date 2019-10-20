package repositories

import (
	. "../model"
	"github.com/jinzhu/gorm"
)

type LivraisonRepository struct {
	db *gorm.DB
}

func NewLivraisonRepository(db *gorm.DB) *LivraisonRepository {
	return &LivraisonRepository{db: db}
}

func (r *LivraisonRepository) Save(m *Livraison) RepositoryResult {
	err := r.db.Save(m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: m}
}
