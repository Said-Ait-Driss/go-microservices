package main

import (
	"net/http"
	"os"
	"os/signal"
	httpHandler "product-service/pkg/transport/http"
	"product-service/pkg/transport/nats/consumer"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	httpHandler.ProductHandler(router)

	httpHandler.OfferHandler(router)

	httpHandler.CategoryHandler(router)

	go consumer.LoadSubscriptions()
	go func() {
		http.ListenAndServe(":3050", router)
	}()

	// Wait for termination signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan


}
