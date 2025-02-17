package mapper

import (
	entity "product-service/pkg/entities"
	"product-service/pkg/transport/http/request"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapUpdateProductRequest(req request.UpdateProductRequest) entity.Product {
	product := entity.Product{
		ID:             req.ID,
		Name:           req.Name,
		Original_price: req.Original_price,
		Store:          req.Store,
		Image:          req.Image,
		Images:         req.Images,
		Quantity:       req.Quantity,
		Current_offer:  req.Current_offer,
		Updated_at:     time.Now(),
	}

	return product
}

func MapCreateProductRequest(req request.CreateProductRequest) entity.Product {
	product := entity.Product{
		ID:             primitive.NewObjectID(),
		Name:           req.Name,
		Original_price: req.Original_price,
		Store:          req.Store,
		Image:          req.Image,
		Images:         req.Images,
		Quantity:       req.Quantity,
		Current_offer:  req.Current_offer,
		Created_at:     time.Now(),
		Updated_at:     time.Now(),
	}

	return product
}
