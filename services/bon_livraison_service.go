package services

import (
	"fmt"
	"log"

	"../dtos"
	. "../model"
	. "../repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateBonLivraison(m *BonLivraison, r BonLivraisonRepository) dtos.Response {
	uuidResult, err := uuid.NewUUID()

	if err != nil {
		log.Fatalln(err)
	}

	m.ID = uint(uuidResult.ID())

	operationResult := r.Save(m)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*BonLivraison)

	return dtos.Response{Success: true, Data: data}
}

func FindAllBonLivraisons(r BonLivraisonRepository) dtos.Response {
	operationResult := r.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*BonLivraisons)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneBonLivraisonById(id uint, r BonLivraisonRepository) dtos.Response {
	operationResult := r.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*BonLivraison)

	return dtos.Response{Success: true, Data: data}
}

func UpdateBonLivraisonById(id uint, m BonLivraison, r BonLivraisonRepository) dtos.Response {
	existingBonLivraisonResponse := FindOneBonLivraisonById(id, r)

	if !existingBonLivraisonResponse.Success {
		return existingBonLivraisonResponse
	}

	existingBonLivraison := existingBonLivraisonResponse.Data.(*BonLivraison)

	existingBonLivraison.PrixAchat = m.PrixAchat
	existingBonLivraison.Quantite = m.Quantite

	operationResult := r.Save(existingBonLivraison)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneBonLivraisonById(id uint, r BonLivraisonRepository) dtos.Response {
	operationResult := r.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteBonLivraisonByIds(multiId *dtos.MultiID, r BonLivraisonRepository) dtos.Response {
	operationResult := r.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func PaginationBonLivraison(r BonLivraisonRepository, c *gin.Context, p *dtos.Pagination) dtos.Response {
	operationResult, totalPages := r.Pagination(p)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*dtos.Pagination)

	// get current url path
	urlPath := c.Request.URL.Path

	// search query params
	searchQueryParams := ""

	for _, search := range p.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	// set first & last page pagination response
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, p.Limit, 0, p.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, p.Limit, totalPages, p.Sort) + searchQueryParams

	if data.Page > 0 {
		// set previous page pagination response
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, p.Limit, data.Page-1, p.Sort) + searchQueryParams
	}

	if data.Page < totalPages {
		// set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, p.Limit, data.Page+1, p.Sort) + searchQueryParams
	}

	if data.Page > totalPages {
		// reset previous page
		data.PreviousPage = ""
	}

	return dtos.Response{Success: true, Data: data}
}
