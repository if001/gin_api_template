package usecase

import (
	"gin_api_template/domain/repository"
	"gin_api_template/domain/model"
)

type CategoryUseCase interface {
	CategoryUseCase() (*[]model.Category, error)
}

type categoryUseCase struct {
	CategoryRepo repository.CategoryRepository
}

func NewCategoryUseCase(cr repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		CategoryRepo: cr,
	}
}

func (u *categoryUseCase) CategoryUseCase() (*[]model.Category, error) {
	categories, err := u.CategoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, err
}
