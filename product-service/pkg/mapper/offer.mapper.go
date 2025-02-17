package mapper

import (
	entity "product-service/pkg/entities"
	"product-service/pkg/transport/http/request"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapCreateOfferRequest(req request.CreateOfferRequest) entity.Offer {
	offer := entity.Offer{
		ID:         primitive.NewObjectID(),
		Price:      req.Price,
		Discount:   req.Discount,
		Start_date: req.Start_date,
		End_date:   req.End_date,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	return offer
}

func MapUpdateOfferRequest(req request.UpdateOfferRequest) entity.Offer {
	offer := entity.Offer{
		ID:         req.ID,
		Price:      req.Price,
		Discount:   req.Discount,
		Start_date: req.Start_date,
		End_date:   req.End_date,
		Updated_at: time.Now(),
	}

	return offer
}
