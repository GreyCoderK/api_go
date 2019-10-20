package repositories

import (
	. "../model"
	"github.com/jinzhu/gorm"
)

type FactureRepository struct {
	db *gorm.DB
}

func NewFactureRepository(db *gorm.DB) *FactureRepository {
	return &FactureRepository{db: db}
}

func (r *FactureRepository) Save(m *Fonction) RepositoryResult {
	err := r.db.Save(m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: m}
}
