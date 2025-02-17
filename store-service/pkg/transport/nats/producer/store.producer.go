package producer

import (
	"encoding/json"
	"log"
	"time"

	request "store-service/pkg/transport/http/request"
	"store-service/pkg/transport/nats"
)

const updateProductsStoreSubject = "updateProductsStore"

func UpdateProductsStore(storeInfo request.StoreInfo) error {
	nc, err := nats.GetNatsInstance()
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	storeInfoJSON, err := json.Marshal(storeInfo)
	if err != nil {
		return err
	}

	response, err := nc.Request(updateProductsStoreSubject, []byte(storeInfoJSON), 500*time.Millisecond)
	if err != nil {
		return err
	}

	// Process the received product information
	log.Printf("Received product information: %s", response.Data)
	return nil
}