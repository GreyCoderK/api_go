package repositories

import (
	"fmt"
	"math"
	"strings"

	. "../dtos"
	. "../model"
	"github.com/jinzhu/gorm"
)

type BonLivraisonRepository struct {
	db *gorm.DB
}

func NewBonLivraisonRepository(db *gorm.DB) *BonLivraisonRepository {
	return &BonLivraisonRepository{db: db}
}

func (r *BonLivraisonRepository) Save(m *BonLivraison) RepositoryResult {
	err := r.db.Save(m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: m}
}

func (r *BonLivraisonRepository) FindAll() RepositoryResult {
	var m BonLivraisons

	err := r.db.Order("created_at").Find(&m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &m}
}

func (r *BonLivraisonRepository) FindOneById(id uint) RepositoryResult {
	var m BonLivraison

	err := r.db.Where(&BonLivraison{Model: gorm.Model{ID: id}}).Take(&m).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &m}
}

func (r *BonLivraisonRepository) DeleteOneById(id uint) RepositoryResult {
	err := r.db.Delete(&BonLivraison{Model: gorm.Model{ID: id}}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *BonLivraisonRepository) DeleteByIds(ids *[]uint) RepositoryResult {
	err := r.db.Where("ID IN (?)", *ids).Delete(&BonLivraison{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *BonLivraisonRepository) Pagination(pagination *Pagination) (RepositoryResult, int) {
	var m BonLivraisons

	totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	find := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	// generate where query
	searchs := pagination.Searchs

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break

			}
		}
	}

	find = find.Find(&m)

	// has error find data
	errFind := find.Error

	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = m

	// count all data
	errCount := r.db.Model(&BonLivraison{}).Count(&totalRows).Error

	if errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}

	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > totalRows {
		// set to row with total rows
		toRow = totalRows
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}
