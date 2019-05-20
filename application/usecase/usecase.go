package usecase

import "gin_api_template/infrastructure"

type UseCase struct {
	BookUseCase
	CategoryUseCase
}

func NewUseCase() UseCase {
	r := infrastructure.NewRepository()
	bu := NewBookUseCase(r.Br)
	cu := NewCategoryUseCase(r.Cr)
	return UseCase{bu , cu}
}