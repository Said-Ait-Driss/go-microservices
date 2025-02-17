package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string             `json:"name" validate:"required"`
	Original_price float64            `json:"original_price" validate:"required,numeric"`
	Store          Store              `json:"store" validate:"required"`
	Image          string             `json:"image" validate:"required"`
	Images         []string           `json:"images"`
	Quantity       int                `json:"quantity" validate:"required"`
	Current_offer  Current_offer      `json:"current_offer" validate:"required"`
	Category_id    string             `json:"category_id" validate:"required"`
	Created_at     time.Time          `json:"created_at"`
	Updated_at     time.Time          `json:"updated_at"`
}
