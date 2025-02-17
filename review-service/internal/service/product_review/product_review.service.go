package service

import (
	repository "review-service/internal/repository/product_review"
	entity "review-service/pkg/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var repo *repository.ProductReviewsRepository

func init() {

	repo = repository.NewProductReviewsRepository()
}

type ProductReviewsService interface {
	GetProductReviews(id primitive.ObjectID) ([]entity.ProductReviews, error)
	CreateProductReview(product_review entity.ProductReviews) ([]entity.ProductReviews, error)
}

// implementations of product service
type ProductReviewsServiceImpl struct{}

func (service *ProductReviewsServiceImpl) GetProductReviews(id primitive.ObjectID) ([]entity.ProductReviews, error) {

	reviews, err := repo.GetProductReviews(id)

	if err != nil {
		return []entity.ProductReviews{}, err
	}

	return reviews, nil
}

func (service *ProductReviewsServiceImpl) CreateProductReview(product_review entity.ProductReviews) ([]entity.ProductReviews, error) {

	product_id := product_review.Product_id

	err := repo.CreateProductReview(product_review)

	if err != nil {
		return []entity.ProductReviews{}, err
	}

	return service.GetProductReviews(product_id)
}
