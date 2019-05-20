package handler

import (
	"github.com/gin-gonic/gin"
	"gin_api_template/application/usecase"
)

type CategoryHandler interface {
	GetCategories(c *gin.Context)
}

type categoryHandler struct {
	CategoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler(c usecase.CategoryUseCase) CategoryHandler {
	return &categoryHandler{
		CategoryUseCase: c,
	}
}

func (u *categoryHandler) GetCategories(c *gin.Context) {
	categories, err := u.CategoryUseCase.CategoryUseCase()
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, categories)
}
