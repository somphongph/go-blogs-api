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

func NewHandler(store storer) *Handler {
	return &Handler{store: store}
}

func (h *Handler) CreateHandler(c echo.Context) error {
	cr := createRequest{}
	if err := c.Bind(&cr); err != nil {
		return c.JSON(http.StatusBadRequest, response.Err{Message: err.Error()})
	}

	// Bind object
	blog := &Blog{
		Id:        primitive.NewObjectID(),
		Title:     cr.Title,
		Content:   cr.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := h.store.Add(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "test")

		return err
	}

	return c.JSON(http.StatusCreated, blog)
}
