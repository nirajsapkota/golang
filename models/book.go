package models

type Book struct {
	ID     int
	Title  string
	Author string
}

type CreateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookRequest struct {
	ID     *int   `json:"id" binding:"required"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type DeleteBookRequest struct {
	ID *int `json:"id" binding:"required"`
}
