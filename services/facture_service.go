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

func CreateFacture(m *Facture, r FactureRepository) dtos.Response {
	uuidResult, err := uuid.NewUUID()

	if err != nil {
		log.Fatalln(err)
	}

	m.ID = uint(uuidResult.ID())

	operationResult := r.Save(m)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*Facture)

	return dtos.Response{Success: true, Data: data}
}

func FindAllFactures(r FactureRepository) dtos.Response {
	operationResult := r.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*Factures)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneFactureById(id uint, r FactureRepository) dtos.Response {
	operationResult := r.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*Facture)

	return dtos.Response{Success: true, Data: data}
}

func UpdateFactureById(id uint, m Facture, r FactureRepository) dtos.Response {
	existingFactureResponse := FindOneFactureById(id, r)

	if !existingFactureResponse.Success {
		return existingFactureResponse
	}

	existingFacture := existingFactureResponse.Data.(*Facture)

	existingFacture.Date = m.Date
	existingFacture.Montant = m.Montant

	operationResult := r.Save(existingFacture)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneFactureById(id uint, r FactureRepository) dtos.Response {
	operationResult := r.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteFactureByIds(multiId *dtos.MultiID, r FactureRepository) dtos.Response {
	operationResult := r.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func PaginationFacture(r FactureRepository, c *gin.Context, p *dtos.Pagination) dtos.Response {
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
