package endpoint

import (
	"context"
	service "product-service/internal/service/category"
	entity "product-service/pkg/entities"
	"product-service/pkg/mapper"
	request "product-service/pkg/transport/http/request"
	response "product-service/pkg/transport/http/response"

	"github.com/go-kit/kit/endpoint"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeCategoryEndpoints(s service.CategoryService) Endpoints {
	return Endpoints{
		CreateCategoryEndpoint: MakeCreateCategoryEndpoint(s),
		GetCategoryEndpoint:    MakeGetCategoryEndpoit(s),
		UpdateCategoryEndpoint: MakeUpdateCategoryEndpoint(s),
		DeleteCategoryEndpoint: MakeDeleteCategoryEndpoint(s),
		GetCategoriesEndpoit:   MakeGetCategoriesEndpoit(s),
	}
}

func MakeCreateCategoryEndpoint(srv service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.CreateCategoryRequest)

		// Validate the request
		if err := req.Validate(); err != nil {
			return response.CreateCategoryResponse{Category: entity.Category{}, Err: err}, nil
		}

		// Map req to entity.Category
		category := mapper.MapCreateCategoryRequest(req)
		_category, err := srv.CreateCategory(category)

		return response.CreateCategoryResponse{Category: _category, Err: err}, nil
	}
}

func MakeGetCategoriesEndpoit(srv service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetCategoriesRequest)

		categories, err := srv.GetCategories(req.Page, req.PageSize)
		return response.GetCategoriesResponse{Categories: categories, Err: err}, nil
	}
}

func MakeGetCategoryEndpoit(srv service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetCategoryRequest)

		// Parse the string ID into primitive.ObjectID
		category_id, err := primitive.ObjectIDFromHex(req.ID)
		if err != nil {
			return response.GetCategoryResponse{Category: entity.Category{}, Err: err}, err
		}

		category, err := srv.GetCategory(category_id)
		return response.GetCategoryResponse{Category: category, Err: err}, nil
	}
}

func MakeUpdateCategoryEndpoint(srv service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.UpdateCategoryRequest)

		// Parse the string ID into primitive.ObjectID
		category_id, err := primitive.ObjectIDFromHex(primitive.ObjectID.Hex(req.ID))
		if err != nil {
			return response.UpdateCategoryResponse{Category: entity.Category{}, Err: err}, err
		}

		// Validate the request
		if err := req.Validate(); err != nil {
			return response.UpdateCategoryResponse{Category: entity.Category{}, Err: err}, nil
		}

		// Map req to entity.category
		category := mapper.MapUpdateCategoryRequest(req)

		_category, err := srv.UpdateCategory(category, category_id)

		return response.UpdateCategoryResponse{Category: _category, Err: err}, nil
	}
}

func MakeDeleteCategoryEndpoint(srv service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.DeleteCategoryRequest)

		// Parse the string ID into primitive.ObjectID
		category_id, err := primitive.ObjectIDFromHex(req.ID)
		if err != nil {
			return response.DeletCategoryResponse{Category: entity.Category{}, Err: err}, err
		}

		err = srv.DeleteCategory(category_id)
		return response.DeletCategoryResponse{Category: entity.Category{ID: category_id}, Err: err}, nil
	}
}
