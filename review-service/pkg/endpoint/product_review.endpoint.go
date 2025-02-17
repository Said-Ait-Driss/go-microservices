package endpoint

import (
	"context"
	service "review-service/internal/service/product_review"
	entity "review-service/pkg/entities"
	"review-service/pkg/mapper"
	request "review-service/pkg/transport/http/request"
	response "review-service/pkg/transport/http/response"

	"github.com/go-kit/kit/endpoint"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeProductReviewsEndpoints(s service.ProductReviewsService) Endpoints {
	return Endpoints{
		GetProductReviewsEndpoint:   MakeGetProductReviewsEndpoint(s),
		CreateProductReviewEndpoint: MakeCreateProductReviewEndpoint(s),
	}
}

func MakeGetProductReviewsEndpoint(srv service.ProductReviewsService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {

		req := _request.(request.GetProductReviewsRequest)

		// Parse the string ID into primitive.ObjectID
		product_id, err := primitive.ObjectIDFromHex(req.ID)
		if err != nil {
			return response.GetProductReviewsResponse{ProductReviews: []entity.ProductReviews{}, Err: err}, err
		}

		reviews, err := srv.GetProductReviews(product_id)
		return response.GetProductReviewsResponse{ProductReviews: reviews, Err: err}, nil
	}
}

func MakeCreateProductReviewEndpoint(srv service.ProductReviewsService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.CreateNewProductReviewRequest)

		// Validate the request
		if err := req.Validate(); err != nil {
			return response.CreateNewProductReviewResponse{ProductReviews: []entity.ProductReviews{}, Err: err}, nil
		}

		// Map req to entity.Product
		product_review := mapper.MapCreateProductReviewRequest(req)
		reviews, err := srv.CreateProductReview(product_review)

		return response.CreateNewProductReviewResponse{ProductReviews: reviews, Err: err}, nil
	}
}
