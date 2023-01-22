package post

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/somphongph/go-post-api/post"
)

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) InitDB() {
	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`

	if _, err := h.db.Exec(createTb); err != nil {
		log.Fatal("can't create table", err)
	}

}

func (h *Handler) CreateExpenseHandler(c echo.Context) error {
	e := Post{}

	if err := c.Bind(&e); err != nil {
		return c.JSON(http.StatusBadRequest, post.Err{Message: err.Error()})
	}

	row := h.db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3,)  RETURNING id",
		e.Title,
	)
	err := row.Scan(&e.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, post.Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, e)
}
