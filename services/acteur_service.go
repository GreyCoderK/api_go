package services

import (
	"fmt"
	"log"

	"../dtos"
	. "../helpers"
	. "../model"
	. "../repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateActeur(m *Acteur, r ActeurRepository) dtos.Response {
	uuidResult, err := uuid.NewUUID()

	if err != nil {
		log.Fatalln(err)
	}

	m.ID = uint(uuidResult.ID())

	if !ValideEmail(m.Email) {
		return dtos.Response{Success: false, Message: "L'email entrez n'est pas valide"}
	}

	m.Pwd = HashAndSalt([]byte(m.Pwd))

	operationResult := r.Save(m)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*Acteur)

	return dtos.Response{Success: true, Data: data}
}

func FindAllActeurs(r ActeurRepository) dtos.Response {
	operationResult := r.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*Acteurs)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneActeurById(id uint, r ActeurRepository) dtos.Response {
	operationResult := r.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*Acteur)

	return dtos.Response{Success: true, Data: data}
}

func UpdateActeurById(id uint, m Acteur, r ActeurRepository) dtos.Response {
	existingActeurResponse := FindOneActeurById(id, r)

	if !existingActeurResponse.Success {
		return existingActeurResponse
	}

	existingActeur := existingActeurResponse.Data.(*Acteur)

	existingActeur.Nom = m.Nom
	existingActeur.Prenom = m.Prenom
	existingActeur.Statut = m.Statut

	if ValideEmail(m.Email) {
		existingActeur.Email = m.Email
	} else {
		return dtos.Response{Success: false, Message: "L'email entrez n'est pas valide"}
	}

	existingActeur.Pwd = HashAndSalt([]byte(m.Pwd))

	operationResult := r.Save(existingActeur)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneActeurById(id uint, r ActeurRepository) dtos.Response {
	operationResult := r.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteActeurByIds(multiId *dtos.MultiID, r ActeurRepository) dtos.Response {
	operationResult := r.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func PaginationActeur(r ActeurRepository, c *gin.Context, p *dtos.Pagination) dtos.Response {
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
