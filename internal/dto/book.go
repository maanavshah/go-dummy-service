package dto

type CreateBookRequestDTO struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Desc   string `json:"desc"`
}
