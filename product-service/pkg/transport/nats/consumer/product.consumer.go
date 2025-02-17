package consumer

import (
	"encoding/json"
	"fmt"
	service "product-service/internal/service/product"
	"product-service/pkg/transport/http/request"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var productService *service.ProductServiceImpl

func init() {
	productService = NewProductService()
}

func NewProductService() *service.ProductServiceImpl {
	return &service.ProductServiceImpl{}
}

func processSubtractProductQty(message *nats.Msg) {

	type Payload struct {
		Product_id primitive.ObjectID
		Quantity   int
	}

	pl := &Payload{}

	//  data := string(m.Data)
	json.Unmarshal(message.Data, pl)
	if pl.Product_id == primitive.NilObjectID {
		replyData := fmt.Sprintf("product id is required ")
		message.Respond([]byte(replyData))
		return
	}
	if pl.Quantity == 0 {
		replyData := fmt.Sprintf("subtract quantity must be greather than 0")
		message.Respond([]byte(replyData))
		return
	}

	err := productService.SubtractProductQty(pl.Product_id, pl.Quantity)
	if err != nil {
		replyData := fmt.Sprintf("error occured while trying update product qty %v", err)
		message.Respond([]byte(replyData))
		return
	}
	replyData := fmt.Sprintf("Ack message # %v", pl)
	message.Respond([]byte(replyData))

	fmt.Printf("I got a message %v\n", pl)
}

func processAddProductQty(message *nats.Msg) {
	type Payload struct {
		Product_id primitive.ObjectID
		Quantity   int
	}

	pl := &Payload{}
	json.Unmarshal(message.Data, pl)

	if pl.Product_id == primitive.NilObjectID {
		replyData := fmt.Sprintf("product id is required ")
		message.Respond([]byte(replyData))
		return
	}

	if pl.Quantity == 0 {
		replyData := fmt.Sprintf("subtract quantity must be greather than 0")
		message.Respond([]byte(replyData))
		return
	}

	err := productService.AddProductQty(pl.Product_id, pl.Quantity)
	if err != nil {
		replyData := fmt.Sprintf("error occured while trying update product qty %v", err)
		message.Respond([]byte(replyData))
		return
	}
	replyData := fmt.Sprintf("Ack message # %v", pl)
	message.Respond([]byte(replyData))

	fmt.Printf("I got a message %v\n", pl)

}

func updateProductsStore(message *nats.Msg) {
	var storeInfo request.StoreInfo
	if err := json.Unmarshal(message.Data, &storeInfo); err != nil {
		replyData := fmt.Sprintf("error occured while trying to get message data %v", err)
		message.Respond([]byte(replyData))
		return
	}
	err := productService.UpdateProductsStore(storeInfo)
	if err != nil {
		replyData := fmt.Sprintf("error occured while trying to update products store %v", err)
		message.Respond([]byte(replyData))
		return
	}
	message.Respond([]byte(nil))
}
