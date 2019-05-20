package handler

import (
	"github.com/gin-gonic/gin"
	"gin_api_template/application/usecase"
	"gin_api_template/application/form"
	"net/http"
)

type BookHandler interface {
	GetBooks(c *gin.Context)
	CreateBook(c *gin.Context)
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, books)
}

func (u *bookHandler) CreateBook(c *gin.Context) {
	var bookForm form.BookForm
	if err := c.ShouldBindJSON(&bookForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookForm)
}