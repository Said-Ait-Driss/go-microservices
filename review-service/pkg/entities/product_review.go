package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID         primitive.ObjectID `bson:"_id"`
	User_id    string             `json:"user_id" validate:"required"`
	Username   string             `json:"username" validate:"required"`
	Full_name  string             `json:"full_name" validate:"required"`
	Email      string             `json:"email" validate:"required,email"`
	Comment    string             `json:"comment" validate:"required"`
	Value      float64            `json:"value" validate:"required,numeric,min=0,max=5"`
	Status     string             `json:"status" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

type ProductReviews struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `json:"title" validate:"required"`
	Product_id primitive.ObjectID `bson:"product_id"`
	Reviews    []Review           `json:"reviews"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
