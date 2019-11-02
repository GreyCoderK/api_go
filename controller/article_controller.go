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

func ArticleCreate(c *gin.Context) {
	var m Article
	r := c.MustGet("articlerepo").(*ArticleRepository)
	err := c.ShouldBindJSON(&m)

	if err != nil {
		res := helpers.GenerateValidationResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	code := http.StatusOK

	res := services.CreateArticle(&m, *r)

	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func ArticleHome(c *gin.Context) {
	code := http.StatusOK
	r := c.MustGet("articlerepo").(*ArticleRepository)

	response := services.FindAllArticles(*r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func ArticleShow(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("articlerepo").(*ArticleRepository)

	if err != nil {
		response := dtos.Response{Success: false, Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := http.StatusOK

	response := services.FindOneArticleById(uint(id), *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func ArticleUpdate(c *gin.Context) {
	id, errors := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("articlerepo").(*ArticleRepository)

	if errors != nil {
		response := dtos.Response{Success: false, Message: errors.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var m Article

	err := c.ShouldBindJSON(&m)

	// validation errors
	if err != nil {
		response := helpers.GenerateValidationResponse(err)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.UpdateArticleById(uint(id), m, *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func ArticleDelete(c *gin.Context) {
	id, errors := strconv.ParseUint(c.Param("id"), 10, 32)
	r := c.MustGet("articlerepo").(*ArticleRepository)

	if errors != nil {
		response := dtos.Response{Success: false, Message: errors.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := http.StatusOK

	response := services.DeleteOneArticleById(uint(id), *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func ArticleDeleteMultiple(c *gin.Context) {
	var multiID dtos.MultiID
	r := c.MustGet("articlerepo").(*ArticleRepository)

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

	response := services.DeleteArticleByIds(&multiID, *r)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}

func ArticlePagination(c *gin.Context) {
	code := http.StatusOK

	r := c.MustGet("articlerepo").(*ArticleRepository)
	pagination := helpers.GeneratePaginationRequest(c)

	response := services.PaginationArticle(*r, c, pagination)

	if !response.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, response)
}
