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

func CategorieArticleCreate(c *gin.Context) {
	var m CategorieArticle
	r := c.MustGet("categoriearticlerepo").(*CategorieArticleRepository)
	err := c.ShouldBindJSON(&m)

	if err != nil {
		res := helpers.GenerateValidationResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	code := http.StatusOK

	res := services.CreateCategorieArticle(&m, *r)

	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func CategorieArticleHome(c *gin.Context) {
	code := http.StatusOK
	r := c.MustGet("categoriearticlerepo").(*CategorieArticleRepository)

	response := services.FindAllCategorieArticles(*r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func CategorieArticleShow(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("categoriearticlerepo").(*CategorieArticleRepository)

	if err != nil {
		response := dtos.Response{Success: false, Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := http.StatusOK

	response := services.FindOneCategorieArticleById(uint(id), *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func CategorieArticleUpdate(c *gin.Context) {
	id, errors := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("categoriearticlerepo").(*CategorieArticleRepository)

	if errors != nil {
		response := dtos.Response{Success: false, Message: errors.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var m CategorieArticle

	err := c.ShouldBindJSON(&m)

	// validation errors
	if err != nil {
		response := helpers.GenerateValidationResponse(err)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.UpdateCategorieArticleById(uint(id), m, *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func CategorieArticleDelete(c *gin.Context) {
	id, errors := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("categoriearticlerepo").(*CategorieArticleRepository)

	if errors != nil {
		response := dtos.Response{Success: false, Message: errors.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := http.StatusOK

	response := services.DeleteOneCategorieArticleById(uint(id), *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func CategorieArticleDeleteMultiple(c *gin.Context) {
	var multiID dtos.MultiID
	r := c.MustGet("categoriearticlerepo").(*CategorieArticleRepository)

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

	response := services.DeleteCategorieArticleByIds(&multiID, *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func CategorieArticlePagination(c *gin.Context) {
	code := http.StatusOK

	r := c.MustGet("categoriearticlerepo").(*CategorieArticleRepository)
	pagination := helpers.GeneratePaginationRequest(c)

	response := services.PaginationCategorieArticle(*r, c, pagination)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}
