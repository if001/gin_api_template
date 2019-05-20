package usecase

import (
	"gin_api_template/domain/model"
	"gin_api_template/domain/repository"
)

type BookUseCase interface {
	GetBooksUseCase() (*[]model.Book, error)
}

type bookUseCase struct {
	BookRep repository.BookRepository
}

func NewBookUseCase(br repository.BookRepository) BookUseCase {
	return &bookUseCase{
		BookRep: br,
	}
}

func (u *bookUseCase) GetBooksUseCase() (*[]model.Book, error) {
	books, err := u.BookRep.GetBooks()
	if err != nil {
		return nil, err
	}
	return books, err
}
