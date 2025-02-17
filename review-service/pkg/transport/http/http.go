package httpHandler

import (
	productReviewCtr "review-service/internal/controller/product_review"
	storeReviewCtr "review-service/internal/controller/store_review"

	"github.com/gofiber/fiber/v2"
)

func ProductReviewsHandler(app *fiber.App) {

	app.Get("/reviews/product/:product_id", productReviewCtr.GetProductReviews)

	// add product review
	app.Post("/reviews/product", productReviewCtr.CreateNewProductReview)
}

func StoreReviewsHandler(app *fiber.App) {

	// get reviews of store
	app.Get("/reviews/store/:store_id", storeReviewCtr.GetStoreReviews)

	// add store review
	app.Post("/create-reviews/store", storeReviewCtr.CreateNewStoreReview)

	// get deleted reviews by store
	app.Get("/reviews/delet-by-store/:store_id", storeReviewCtr.GetDeletedReviewsByStore)

	// get deleted reviews by user
	app.Get("/reviews/delet-by-user/:user_id", storeReviewCtr.GetDeletedReviewsByUser)

	// get reviews of list of stores
	app.Get("/reviews/stores/:store_ids", storeReviewCtr.GetReviewsOfListOfStores)

	// change review status to = deleted, pending_to_deleted_by_user, pending_to_deleted_by_store
	app.Put("/review/change-status", storeReviewCtr.ChangeReviewStatus)
}
