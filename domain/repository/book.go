package repository

import "gin_api_template/domain/model"

type BookRepository interface {
	GetBooks() (*[]model.Book, error)
}

