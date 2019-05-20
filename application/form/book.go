package form

type BookForm struct {
	Name   string `json:"name" binding:"required"`
	Author string `json:"author"`
}
