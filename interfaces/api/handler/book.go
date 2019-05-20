package handler

import (
	"github.com/gin-gonic/gin"
	"gin_api_template/application/usecase"
)

type BookHandler interface {
	GetBooks(c *gin.Context)
}

type bookHandler struct {
	BookUseCase usecase.BookUseCase
}

func NewBookHandler(u usecase.BookUseCase) BookHandler {
	return &bookHandler{
		BookUseCase: u,
	}
}

func (u *bookHandler) GetBooks(c *gin.Context) {
	books, err := u.BookUseCase.GetBooksUseCase()
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, books)
}
