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

func LivraisonCreate(c *gin.Context) {
	var m Livraison
	r := c.MustGet("Livraisonrepo").(*LivraisonRepository)
	err := c.ShouldBindJSON(&m)

	if err != nil {
		res := helpers.GenerateValidationResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	code := http.StatusOK

	res := services.CreateLivraison(&m, *r)

	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func LivraisonHome(c *gin.Context) {
	code := http.StatusOK
	r := c.MustGet("Livraisonrepo").(*LivraisonRepository)

	response := services.FindAllLivraisons(*r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func LivraisonShow(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("Livraisonrepo").(*LivraisonRepository)

	if err != nil {
		response := dtos.Response{Success: false, Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := http.StatusOK

	response := services.FindOneLivraisonById(uint(id), *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func LivraisonUpdate(c *gin.Context) {
	id, errors := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("Livraisonrepo").(*LivraisonRepository)

	if errors != nil {
		response := dtos.Response{Success: false, Message: errors.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var m Livraison

	err := c.ShouldBindJSON(&m)

	// validation errors
	if err != nil {
		response := helpers.GenerateValidationResponse(err)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.UpdateLivraisonById(uint(id), m, *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func LivraisonDelete(c *gin.Context) {
	id, errors := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("Livraisonrepo").(*LivraisonRepository)

	if errors != nil {
		response := dtos.Response{Success: false, Message: errors.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := http.StatusOK

	response := services.DeleteOneLivraisonById(uint(id), *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func LivraisonDeleteMultiple(c *gin.Context) {
	var multiID dtos.MultiID
	r := c.MustGet("Livraisonrepo").(*LivraisonRepository)

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

	response := services.DeleteLivraisonByIds(&multiID, *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func LivraisonPagination(c *gin.Context) {
	code := http.StatusOK

	r := c.MustGet("Livraisonrepo").(*LivraisonRepository)
	pagination := helpers.GeneratePaginationRequest(c)

	response := services.PaginationLivraison(*r, c, pagination)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}
