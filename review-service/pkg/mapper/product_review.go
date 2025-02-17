package mapper

import (
	entity "review-service/pkg/entities"
	"review-service/pkg/transport/http/request"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapCreateProductReviewRequest(req request.CreateNewProductReviewRequest) entity.ProductReviews {
	product_review := entity.ProductReviews{
		ID:         primitive.NewObjectID(),
		Title:      req.Product_title,
		Product_id: req.Product_id,
		Reviews:    []entity.Review{req.Review},
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	return product_review
}
