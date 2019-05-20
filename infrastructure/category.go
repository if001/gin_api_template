package infrastructure

import (
	"github.com/jinzhu/gorm"
	"gin_api_template/domain/model"
	"gin_api_template/domain/repository"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepository{ DB : db }
}

var categories []model.Category
func (c *categoryRepository) GetCategories() (*[]model.Category, error) {
	err  := c.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return &categories, nil
}