package repositories

import (
	. "../model"
	"github.com/jinzhu/gorm"
)

type StructureRepository struct {
	db *gorm.DB
}

func NewStructureRepository(db *gorm.DB) *StructureRepository {
	return &StructureRepository{db: db}
}

func (r *StructureRepository) Save(m *Structure) RepositoryResult {
	err := r.db.Save(m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: m}
}

func (r *StructureRepository) FindAll() RepositoryResult {
	var m Structure

	err := r.db.Find(&m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &m}
}

func (r *StructureRepository) FindOneById(id uint) RepositoryResult {
	var m Structure

	err := r.db.Where(&Structure{Model: gorm.Model{ID: id}}).Take(&m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &m}
}

func (r *StructureRepository) DeleteOneById(id uint) RepositoryResult {
	err := r.db.Delete(&Structure{Model: gorm.Model{ID: id}}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *StructureRepository) DeleteByIds(ids *[]uint) RepositoryResult {
	err := r.db.Where("ID IN (?)", *ids).Delete(&Structure{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
