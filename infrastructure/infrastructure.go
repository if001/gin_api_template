package infrastructure

import (
	"gin_api_template/domain/repository"
	"gin_api_template/infrastructure/mysql_conn"
)

type Repository struct {
	Br repository.BookRepository
	Cr repository.CategoryRepository
}

func NewRepository() Repository {
	db := mysql_conn.GetDBConn()

	br := NewBookRepository(db)
	cr := NewCategoryRepository(db)
	return Repository{Br: br, Cr: cr}
}
