package post

import "database/sql"

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type Handler struct {
	db *sql.DB
}

type Err struct {
	Message string `json:"message"`
}
