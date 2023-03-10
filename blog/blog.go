package blog

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type Handler struct {
	store storer
}

type storer interface {
	GetById(string) (Blog, error)
	GetList() ([]Blog, error)
	Add(*Blog) error
	Update(*Blog) error
	Delete(string) error
}

func New(store storer) *Handler {
	return &Handler{store: store}
}
