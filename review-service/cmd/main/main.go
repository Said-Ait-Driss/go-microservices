package main

import (
	httpHandler "review-service/pkg/transport/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

func main() {

	app := fiber.New()

	httpHandler.ProductReviewsHandler(app)
	httpHandler.StoreReviewsHandler(app)

	app.Listen(":3054")

}
