package controller

import (
	"net/http"
	"strconv"

	"../dtos"
	"../helpers"
	. "../model"
	. "../repositories"
	"../services"
	"github.com/gin-gonic/gin"
)

func FactureCreate(c *gin.Context) {
	var m Facture
	r := c.MustGet("facturerepo").(*FactureRepository)
	err := c.ShouldBindJSON(&m)

	if err != nil {
		res := helpers.GenerateValidationResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	code := http.StatusOK

	res := services.CreateFacture(&m, *r)

	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func FactureHome(c *gin.Context) {
	code := http.StatusOK
	r := c.MustGet("facturerepo").(*FactureRepository)

	response := services.FindAllFactures(*r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func FactureShow(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("facturerepo").(*FactureRepository)

	if err != nil {
		response := dtos.Response{Success: false, Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := http.StatusOK

	response := services.FindOneFactureById(uint(id), *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func FactureUpdate(c *gin.Context) {
	id, errors := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("facturerepo").(*FactureRepository)

	if errors != nil {
		response := dtos.Response{Success: false, Message: errors.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var m Facture

	err := c.ShouldBindJSON(&m)

	// validation errors
	if err != nil {
		response := helpers.GenerateValidationResponse(err)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.UpdateFactureById(uint(id), m, *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func FactureDelete(c *gin.Context) {
	id, errors := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("facturerepo").(*FactureRepository)

	if errors != nil {
		response := dtos.Response{Success: false, Message: errors.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := http.StatusOK

	response := services.DeleteOneFactureById(uint(id), *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func FactureDeleteMultiple(c *gin.Context) {
	var multiID dtos.MultiID
	r := c.MustGet("facturerepo").(*FactureRepository)

	err := c.ShouldBindJSON(&multiID)

	// validation errors
	if err != nil {
		response := helpers.GenerateValidationResponse(err)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	if len(multiID.Ids) == 0 {
		response := dtos.Response{Success: false, Message: "IDs ne peut-être vide"}

		c.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.DeleteFactureByIds(&multiID, *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func FacturePagination(c *gin.Context) {
	code := http.StatusOK

	r := c.MustGet("facturerepo").(*FactureRepository)
	pagination := helpers.GeneratePaginationRequest(c)

	response := services.PaginationFacture(*r, c, pagination)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}
