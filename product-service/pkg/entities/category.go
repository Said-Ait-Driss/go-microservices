package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `json:"title" validate:"required"`
	Image      string             `json:"image" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
