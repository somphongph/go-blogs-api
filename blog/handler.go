package blog

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/somphongph/go-blogs-api/response"
)

func (h *Handler) CreateHandlerXXXX(c echo.Context) error {
	e := Blog{}

	if err := c.Bind(&e); err != nil {
		return c.JSON(http.StatusBadRequest, response.Err{Message: err.Error()})
	}

	return c.JSON(http.StatusInternalServerError, response.Err{Message: ""})
}
