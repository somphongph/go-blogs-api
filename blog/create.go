package blog

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) CreateExpenseHandler(c echo.Context) error {
	e := Blog{}

	if err := c.Bind(&e); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := h.db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3,)  RETURNING id",
		e.Title,
	)
	err := row.Scan(&e.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, e)
}
