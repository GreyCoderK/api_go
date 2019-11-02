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

func CreateLivraison(m *Livraison, r LivraisonRepository) dtos.Response {
	uuidResult, err := uuid.NewUUID()

	if err != nil {
		log.Fatalln(err)
	}

	m.ID = uint(uuidResult.ID())

	operationResult := r.Save(m)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*Livraison)

	return dtos.Response{Success: true, Data: data}
}

func FindAllLivraisons(r LivraisonRepository) dtos.Response {
	operationResult := r.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*Livraisons)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneLivraisonById(id uint, r LivraisonRepository) dtos.Response {
	operationResult := r.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*Livraison)

	return dtos.Response{Success: true, Data: data}
}

func UpdateLivraisonById(id uint, m Livraison, r LivraisonRepository) dtos.Response {
	existingLivraisonResponse := FindOneLivraisonById(id, r)

	if !existingLivraisonResponse.Success {
		return existingLivraisonResponse
	}

	existingLivraison := existingLivraisonResponse.Data.(*Livraison)

	existingLivraison.Date = m.Date
	existingLivraison.Montant = m.Montant

	operationResult := r.Save(existingLivraison)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneLivraisonById(id uint, r LivraisonRepository) dtos.Response {
	operationResult := r.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteLivraisonByIds(multiId *dtos.MultiID, r LivraisonRepository) dtos.Response {
	operationResult := r.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func PaginationLivraison(r LivraisonRepository, c *gin.Context, p *dtos.Pagination) dtos.Response {
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
