package handler

import (
	"gin_api_template/application/usecase"
)

type ApiHandler interface {
	BookHandler
	CategoryHandler
}

type apiHandler struct {
	BookHandler
	CategoryHandler
}

func NewApiHandler() ApiHandler {
	app := usecase.NewUseCase()
	bh := NewBookHandler(app.BookUseCase)
	ch := NewCategoryHandler(app.CategoryUseCase)

	return &apiHandler{bh, ch}
}