package mapper

import (
	entity "review-service/pkg/entities"
	"review-service/pkg/transport/http/request"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapCreateStoreReviewRequest(req request.CreateNewStoreReviewRequest) entity.StoreReviews {
	store_review := entity.StoreReviews{
		ID:         primitive.NewObjectID(),
		Title:      req.Store_title,
		Store_id:   req.Store_id,
		Reviews:    []entity.Review{MapReviewToReviewArray(req)},
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	return store_review
}

func MapReviewToReviewArray(req request.CreateNewStoreReviewRequest) entity.Review {
	review := entity.Review{
		ID:         primitive.NewObjectID(),
		User_id:    req.Review.User_id,
		Username:   req.Review.Username,
		Full_name:  req.Review.Full_name,
		Email:      req.Review.Email,
		Comment:    req.Review.Comment,
		Value:      req.Review.Value,
		Status:     req.Review.Status,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	return review
}
