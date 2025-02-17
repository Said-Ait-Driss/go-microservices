package controller

import (
	"context"
	"net/http"
	service "review-service/internal/service/product_review"
	"review-service/pkg/endpoint"
	"review-service/pkg/transport/http/request"
	"review-service/pkg/transport/http/response"

	"github.com/gofiber/fiber/v2"
)

var productReviewsEndpoints endpoint.Endpoints

func init() {

	productReviewsService := &service.ProductReviewsServiceImpl{}

	productReviewsEndpoints = endpoint.MakeProductReviewsEndpoints(productReviewsService)

}

// ConvertFiberToContext converts Fiber context to context.Context
func ConvertFiberToContext(ctx *fiber.Ctx) context.Context {
	// Create a new context
	return context.Background()
}

func GetProductReviews(ctx *fiber.Ctx) error {
	// Convert Fiber context to context.Context
	context := ConvertFiberToContext(ctx)

	product_id := ctx.Params("product_id")

	_request := request.GetProductReviewsRequest{ID: product_id}
	_response, _ := productReviewsEndpoints.GetProductReviewsEndpoint(context, _request)

	res := _response.(response.GetProductReviewsResponse)

	if res.Err != nil {
		response := fiber.Map{"errors": res.Err.Error()}
		ctx.Status(http.StatusBadRequest).JSON(response)
		return nil
	}
	response := fiber.Map{"reviews": res.ProductReviews}
	ctx.Status(fiber.StatusOK).JSON(response)
	return nil
}

func CreateNewProductReview(ctx *fiber.Ctx) error {
	context := ConvertFiberToContext(ctx)

	var request request.CreateNewProductReviewRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := fiber.Map{"errors": err.Error()}
		ctx.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	_response, _ := productReviewsEndpoints.CreateProductReviewEndpoint(context, request)

	res := _response.(response.CreateNewProductReviewResponse)

	if res.Err != nil {
		response := fiber.Map{"errors": res.Err.Error()}
		ctx.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := fiber.Map{"reviews": res.ProductReviews}
	ctx.Status(http.StatusBadRequest).JSON(response)
	return nil
}
