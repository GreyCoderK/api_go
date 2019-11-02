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

func BonLivraisonCreate(c *gin.Context) {
	var m BonLivraison
	r := c.MustGet("bonlivraisonrepo").(*BonLivraisonRepository)
	err := c.ShouldBindJSON(&m)

	if err != nil {
		res := helpers.GenerateValidationResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	code := http.StatusOK

	res := services.CreateBonLivraison(&m, *r)

	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func BonLivraisonHome(c *gin.Context) {
	code := http.StatusOK
	r := c.MustGet("bonlivraisonrepo").(*BonLivraisonRepository)

	response := services.FindAllBonLivraisons(*r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func BonLivraisonShow(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("bonlivraisonrepo").(*BonLivraisonRepository)

	if err != nil {
		response := dtos.Response{Success: false, Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := http.StatusOK

	response := services.FindOneBonLivraisonById(uint(id), *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func BonLivraisonUpdate(c *gin.Context) {
	id, errors := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("bonlivraisonrepo").(*BonLivraisonRepository)

	if errors != nil {
		response := dtos.Response{Success: false, Message: errors.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var m BonLivraison

	err := c.ShouldBindJSON(&m)

	// validation errors
	if err != nil {
		response := helpers.GenerateValidationResponse(err)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.UpdateBonLivraisonById(uint(id), m, *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func BonLivraisonDelete(c *gin.Context) {
	id, errors := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("bonlivraisonrepo").(*BonLivraisonRepository)

	if errors != nil {
		response := dtos.Response{Success: false, Message: errors.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := http.StatusOK

	response := services.DeleteOneBonLivraisonById(uint(id), *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func BonLivraisonDeleteMultiple(c *gin.Context) {
	var multiID dtos.MultiID
	r := c.MustGet("bonlivraisonrepo").(*BonLivraisonRepository)

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

	response := services.DeleteBonLivraisonByIds(&multiID, *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func BonLivraisonPagination(c *gin.Context) {
	code := http.StatusOK

	r := c.MustGet("bonlivraisonrepo").(*BonLivraisonRepository)
	pagination := helpers.GeneratePaginationRequest(c)

	response := services.PaginationBonLivraison(*r, c, pagination)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}
