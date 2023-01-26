package blog

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/somphongph/go-blogs-api/response"
)

func (h *Handler) CreateHandler(c echo.Context) error {

	return c.JSON(http.StatusInternalServerError, response.Err{Message: ""})
}
