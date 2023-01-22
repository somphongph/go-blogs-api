package blog

import "database/sql"

type Blog struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type Handler struct {
	db *sql.DB
}

type Err struct {
	Message string `json:"message"`
}
