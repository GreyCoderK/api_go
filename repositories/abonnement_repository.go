package repositories

import (
	. "../model"
	"github.com/jinzhu/gorm"
)

type AbonnementRepository struct {
	db *gorm.DB
}

func NewAbonnementRepository(db *gorm.DB) *AbonnementRepository {
	return &AbonnementRepository{db: db}
}

func (r *AbonnementRepository) Save(m *Fonction) RepositoryResult {
	err := r.db.Save(m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: m}
}
