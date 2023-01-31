package blog

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type blogResponse struct {
	Id string `json:"id"`
}

func (h *Handler) GetById(c echo.Context) error {
	id := c.Param("id")
	// var blog blogResponse

	// Get data
	blog, err := h.store.GetById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}

	return c.JSON(http.StatusCreated, blog)
}
