package repositories

import (
	. "../model"
	"github.com/jinzhu/gorm"
)

type ActeurRepository struct {
	db *gorm.DB
}

func NewActeurRepository(db *gorm.DB) *ActeurRepository {
	return &ActeurRepository{db: db}
}

func (r *ActeurRepository) Save(m *Fonction) RepositoryResult {
	err := r.db.Save(m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: m}
}
