package repository

import (
	"gin_api_template/domain/model"
)

type CategoryRepository interface {
	GetCategories() (*[]model.Category, error)
}

