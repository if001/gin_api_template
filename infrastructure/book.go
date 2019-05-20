package infrastructure

import (
	"github.com/jinzhu/gorm"
	"gin_api_template/domain/model"
	"gin_api_template/domain/repository"
)

type bookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{ DB : db }
}

func (c *bookRepository) GetBooks() (*[]model.Book, error) {
	books := []model.Book{}
	err := c.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return &books, nil
}
