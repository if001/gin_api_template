package model

import "bookshelf-web-api/domain/service"

type Book struct {
	BaseModel
	Name           string
	AccountId      int64
	AuthorId       int64
	PublishedAt    service.NullTime
	Publisher      string
	StartAt        service.NullTime
	EndAt          service.NullTime
	NextBookID     service.NullInt64
	PrevBookID     service.NullInt64
	DescriptionsId []int64
	// CategoriesId   []int64
}
