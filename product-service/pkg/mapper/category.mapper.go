package mapper

import (
	entity "product-service/pkg/entities"
	"product-service/pkg/transport/http/request"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapCreateCategoryRequest(req request.CreateCategoryRequest) entity.Category {
	category := entity.Category{
		ID:         primitive.NewObjectID(),
		Title:      req.Title,
		Image:      req.Image,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	return category
}

func MapUpdateCategoryRequest(req request.UpdateCategoryRequest) entity.Category {
	category := entity.Category{
		ID:         req.ID,
		Title:      req.Title,
		Image:      req.Image,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	return category
}
