package repositories

import (
	. "../model"
	"github.com/jinzhu/gorm"
)

type CategorieArticleRepository struct {
	db *gorm.DB
}

func NewCategorieArticleRepository(db *gorm.DB) *CategorieArticleRepository {
	return &CategorieArticleRepository{db: db}
}

func (r *CategorieArticleRepository) Save(m *Fonction) RepositoryResult {
	err := r.db.Save(m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: m}
}
