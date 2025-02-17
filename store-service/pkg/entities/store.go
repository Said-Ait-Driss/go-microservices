package entity

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID                primitive.ObjectID  `bson:"_id"`
	Name              string              `json:"name" validate:"required"`
	Description       string              `json:"description" validate:"required"`
	Phone 		      string              `json:"phone" validate:"required"`
	Logo_url          string              `json:"logo_url"`
	Covers            []Covers            `json:"covers"`
	Longitude         string              `json:"longitude" validate:"required"`
	Latitude          string              `json:"latitude" validate:""`
	Client_id         string              `json:"client_id" validate:"required"`
	Adresse           string           	  `json:"adresse" validate:"required"`
	City    	      string          	  `json:"city" validate:"required"`
	Neighborhood      string          	  `json:"neighborhood" validate:"required"`
	Category          Categories          `json:"category" validate:"required"`
	Status            string          	  `json:"status" validate:"required"`
	Open_from          string           	  `json:"open_from" validate:"required"`
	Open_to            string              `json:"open_to" validate:"required"`
	Created_at     	  time.Time           `json:"created_at"`
	Updated_at        time.Time           `json:"updated_at"`
}

type Categories struct {
	ID       	primitive.ObjectID 	`bson:"_id" validate:"required"`
	Code    	string              `json:"code" validate:"required"`
	Libelle    	string              `json:"libelle" validate:"required"`
	Created_at  time.Time           `json:"created_at"`
	Updated_at  time.Time           `json:"updated_at"`
}

type Covers struct {
	ID       	primitive.ObjectID 	`bson:"_id" validate:"required"`
	Url    		string              `json:"url" validate:"required"`
	Order    	string              `json:"order" validate:"required"`
}
