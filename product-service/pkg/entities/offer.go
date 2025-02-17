package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Current_offer struct {
	Offer_id string  `json:"offer_id" validate:"required"`
	Price    float64 `json:"price" validate:"required,numeric"`
	Discount float64 `json:"discount" validate:"required,numeric"`
}

type Offer struct {
	ID         primitive.ObjectID `bson:"_id"`
	Price      float64            `json:"price" validate:"required,numeric"`
	Discount   float64            `json:"discount" validate:"required,numeric"`
	Start_date string             `json:"start_date" validate:"required"`
	End_date   string             `json:"end_date" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
