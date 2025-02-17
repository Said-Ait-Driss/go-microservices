package endpoint

import (
	"context"
	service "product-service/internal/service/product"
	entity "product-service/pkg/entities"
	"product-service/pkg/mapper"
	request "product-service/pkg/transport/http/request"
	response "product-service/pkg/transport/http/response"

	"github.com/go-kit/kit/endpoint"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeProductEndpoints(s service.ProductService) Endpoints {
	return Endpoints{
		GetProductEndpoint:                      MakeGetProductEndpoint(s),
		GetProductsEndpoit:                      MakeGetProductsEndpoint(s),
		CreateProductEndpoint:                   MakeCreateProductEndpoint(s),
		UpdateProductEndpoint:                   MakeUpdateProductEndpoint(s),
		DeleteProductEndpoint:                   MakeDeleteProductEndpoint(s),
		GetProductsCountByStoreEndpoint:         MakeGetProductsCountEndpoint(s),
		GetProductsThatHasOffersEndpoint:        MakeGetProductsThatHasOffersEndpoint(s),
		GetProductsOfStoreThatHasOffersEndpoint: MakeGetProductsOfStoreThatHasOffers(s),
		GetStoreProductsFilterEndpoint:          MakeGetStoreProductsFilterEndpoint(s),
		GetProductsFilterEndpoint:               MakeGetProductsFilterEndpoint(s),
		GetProductsByCategoryEndpoint:           MakeGetProductsByCategoryEndpoint(s),
		GetProductsByStoreEndpoint:              MakeGetProductsByStoreEndpoint(s),
	}
}

/**
 * return endpoint for getting product
 * @func
 * @param {Service} srv product service
 */
func MakeGetProductEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetProductRequest)

		// Parse the string ID into primitive.ObjectID
		product_id, err := primitive.ObjectIDFromHex(req.ID)
		if err != nil {
			return response.GetProductResponse{Product: entity.Product{}, Err: err}, err
		}

		product, err := srv.GetProduct(product_id)
		return response.GetProductResponse{Product: product, Err: err}, nil
	}
}

func MakeGetProductsEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetProductsRequest)

		products, err := srv.GetProducts(req.Page, req.PageSize)
		return response.GetProductsResponse{Products: products, Err: err}, nil
	}
}

func MakeCreateProductEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.CreateProductRequest)

		// Validate the request
		if err := req.Validate(); err != nil {
			return response.CreateProductResponse{Product: entity.Product{}, Err: err}, nil
		}

		// Map req to entity.Product
		product := mapper.MapCreateProductRequest(req)
		_product, err := srv.CreateProduct(product)

		return response.CreateProductResponse{Product: _product, Err: err}, nil
	}
}

func MakeUpdateProductEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.UpdateProductRequest)

		// Parse the string ID into primitive.ObjectID
		product_id, err := primitive.ObjectIDFromHex(primitive.ObjectID.Hex(req.ID))
		if err != nil {
			return response.UpdateProductResponse{Product: entity.Product{}, Err: err}, err
		}

		// Validate the request
		if err := req.Validate(); err != nil {
			return response.UpdateProductResponse{Product: entity.Product{}, Err: err}, nil
		}

		// Map req to entity.Product
		product := mapper.MapUpdateProductRequest(req)

		_product, err := srv.UpdateProduct(product, product_id)

		return response.UpdateProductResponse{Product: _product, Err: err}, nil
	}
}

func MakeDeleteProductEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.DeleteProductRequest)

		// Parse the string ID into primitive.ObjectID
		product_id, err := primitive.ObjectIDFromHex(req.ID)
		if err != nil {
			return response.DeletProductResponse{Product: entity.Product{}, Err: err}, err
		}

		err = srv.DeleteProduct(product_id)
		return response.DeletProductResponse{Product: entity.Product{ID: product_id}, Err: err}, nil
	}
}

// get products count of a store
func MakeGetProductsCountEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetProductsCountByStoreRequest)
		count, err := srv.GetProductsCount(req.Store_id)
		return response.GetProductsCountResponse{Count: count, Err: err}, nil
	}

}

// get products that has offers
func MakeGetProductsThatHasOffersEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetProductsRequest)
		products, err := srv.GetProductsThatHasOffers(req.Page, req.PageSize)
		return response.GetProductsResponse{Products: products, Err: err}, nil
	}
}

// get products of store that has offers
func MakeGetProductsOfStoreThatHasOffers(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetProductsOfStoreThatHasOffersRequest)
		products, err := srv.GetProductsOfStoreThatHasOffers(req.Store_id, req.Page, req.PageSize)
		return response.GetProductsResponse{Products: products, Err: err}, nil
	}
}

// get products filtered
func MakeGetStoreProductsFilterEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetStoreProductsFilterRequest)
		products, err := srv.GetStoreProductsFilter(req.Store_id, req.Name, req.MinPrice, req.MaxPrice, req.StartDate, req.EndDate, req.Page, req.PageSize)
		return response.GetProductsResponse{Products: products, Err: err}, nil
	}
}

func MakeGetProductsFilterEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetProductsFilterRequest)
		products, err := srv.GetProductsFilter(req.Name, req.MinPrice, req.MaxPrice, req.StartDate, req.EndDate, req.Neighborhood, req.StoreTitle, req.Page, req.PageSize)
		return response.GetProductsResponse{Products: products, Err: err}, nil
	}
}

func MakeGetProductsByCategoryEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetProductByCategoryRequest)
		products, err := srv.GetProductsByCategory(req.CategoryId)
		return response.GetProductsResponse{Products: products, Err: err}, nil
	}
}

// MakeGetProductsByStoreEndpoint
func MakeGetProductsByStoreEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetProductsByStoreRequest)
		products, err := srv.GetProductsByStore(req.Store_id)
		return response.GetProductsResponse{Products: products, Err: err}, nil
	}
}
