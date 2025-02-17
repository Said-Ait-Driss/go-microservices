package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoreReviews struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `json:"title" validate:"required"`
	Store_id   string             `json:"store_id"`
	Reviews    []Review           `json:"reviews"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
