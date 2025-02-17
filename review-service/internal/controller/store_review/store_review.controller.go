package controller

import (
	"context"
	"net/http"
	service "review-service/internal/service/store_review"
	"review-service/pkg/endpoint"
	"review-service/pkg/transport/http/request"
	"review-service/pkg/transport/http/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var storeReviewsEndpoints endpoint.Endpoints

func init() {

	storeReviewsService := &service.StoreReviewsServiceImpl{}

	storeReviewsEndpoints = endpoint.MakeStoreReviewsEndpoints(storeReviewsService)

}

// ConvertFiberToContext converts Fiber context to context.Context
func ConvertFiberToContext(ctx *fiber.Ctx) context.Context {
	// Create a new context
	return context.Background()
}

func GetStoreReviews(ctx *fiber.Ctx) error {
	// Convert Fiber context to context.Context
	context := ConvertFiberToContext(ctx)

	store_id := ctx.Params("store_id")

	_request := request.GetStoreReviewsRequest{ID: store_id}
	_response, _ := storeReviewsEndpoints.GetStoreReviewsEndpoint(context, _request)

	res := _response.(response.GetStoreReviewsResponse)

	if res.Err != nil {
		response := fiber.Map{"errors": res.Err.Error()}
		ctx.Status(http.StatusBadRequest).JSON(response)
		return nil
	}
	response := fiber.Map{"reviews": res.StoreReviews}
	ctx.Status(fiber.StatusOK).JSON(response)
	return nil
}

func CreateNewStoreReview(ctx *fiber.Ctx) error {
	context := ConvertFiberToContext(ctx)

	var request request.CreateNewStoreReviewRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := fiber.Map{"errors": err.Error()}
		ctx.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	_response, _ := storeReviewsEndpoints.CreateStoreReviewEndpoint(context, request)

	res := _response.(response.CreateNewStoreReviewResponse)

	if res.Err != nil {
		response := fiber.Map{"errors": res.Err.Error()}
		ctx.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	response := fiber.Map{"reviews": res.StoreReviews}
	ctx.Status(http.StatusBadRequest).JSON(response)
	return nil
}

// deleted reviews by store
func GetDeletedReviewsByStore(ctx *fiber.Ctx) error {
	// Convert Fiber context to context.Context
	context := ConvertFiberToContext(ctx)

	store_id := ctx.Params("store_id")
	_request := request.DeleteReviewByStoreRequest{Store_id: store_id}

	_response, _ := storeReviewsEndpoints.GetDeletedReviewsByStoreEndpoint(context, _request)

	res := _response.(response.GetDeletedReviewsByStoreResponse)

	if res.Err != nil {
		response := fiber.Map{"errors": res.Err.Error()}
		ctx.Status(http.StatusBadRequest).JSON(response)
		return nil
	}
	response := fiber.Map{"reviews": res}
	ctx.Status(fiber.StatusOK).JSON(response)
	return nil
}

// deleted reviews by user
func GetDeletedReviewsByUser(ctx *fiber.Ctx) error {
	// Convert Fiber context to context.Context
	context := ConvertFiberToContext(ctx)

	user_id := ctx.Params("user_id")
	_request := request.DeleteReviewByUserRequest{User_id: user_id}

	_response, _ := storeReviewsEndpoints.GetDeletedReviewsByUserEndpoint(context, _request)

	res := _response.(response.GetDeletedReviewsByUserResponse)

	if res.Err != nil {
		response := fiber.Map{"errors": res.Err.Error()}
		ctx.Status(http.StatusBadRequest).JSON(response)
		return nil
	}
	response := fiber.Map{"reviews": res}
	ctx.Status(fiber.StatusOK).JSON(response)
	return nil
}

func ChangeReviewStatus(ctx *fiber.Ctx) error {
	// Convert Fiber context to context.Context
	context := ConvertFiberToContext(ctx)

	var request request.ChangeReviewStatusRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := fiber.Map{"errors": err.Error()}
		ctx.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	_response, _ := storeReviewsEndpoints.ChangeReviewStatusEndpoint(context, request)

	res := _response.(response.GetSingleStoreReviewsResponse)

	if res.Err != nil {
		response := fiber.Map{"errors": res.Err.Error()}
		ctx.Status(http.StatusBadRequest).JSON(response)
		return nil
	}
	response := fiber.Map{"review": res.StoreReviews}
	ctx.Status(fiber.StatusOK).JSON(response)
	return nil
}

// get reviews of a list of stores
func GetReviewsOfListOfStores(ctx *fiber.Ctx) error {
	// Convert Fiber context to context.Context
	context := ConvertFiberToContext(ctx)

	storeIDsParam := ctx.Params("store_ids")

	storeIDs := strings.Split(storeIDsParam, ",")

	_request := request.GetReviewsOfListOfStoresRequest{Store_ids: storeIDs}

	_response, _ := storeReviewsEndpoints.GetReviewsOfListOfStoresEndpoint(context, _request)

	res := _response.(response.GetStoreReviewsResponse)

	if res.Err != nil {
		response := fiber.Map{"errors": res.Err.Error()}
		ctx.Status(http.StatusBadRequest).JSON(response)
		return nil
	}
	response := fiber.Map{"reviews": res.StoreReviews}
	ctx.Status(fiber.StatusOK).JSON(response)
	return nil
}
