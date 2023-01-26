package blog

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/somphongph/go-blogs-api/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type createResponse struct {
	Id string `json:"id"`
}

func (h *Handler) CreateHandler(c echo.Context) error {
	req := createRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Err{Message: err.Error()})
	}

	// Bind object
	blog := &Blog{
		Id:        primitive.NewObjectID(),
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := h.store.Add(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Err{Message: err.Error()})

		return err
	}

	res := createResponse{
		Id: blog.Id.Hex(),
	}

	return c.JSON(http.StatusCreated, res)
}
