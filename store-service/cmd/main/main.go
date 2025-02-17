package main

import (
	"log"
	"net/http"
	
	httpHandler "store-service/pkg/transport/http"
	storeConsumer "store-service/pkg/transport/nats/consumer"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Attach your handlers to the appropriate route groups
	storeGroup := router.Group("/store")
	categoryGroup := router.Group("/category")

	httpHandler.StoreHandler(storeGroup)
	httpHandler.CategoryHandler(categoryGroup)

	go storeConsumer.LoadSubscriptions()

	// Start the server
	port := ":3053"
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
