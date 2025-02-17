package consumer

import (
	"encoding/json"
	"fmt"
	service "store-service/internal/service/store"
	"store-service/pkg/transport/http/request"

	"github.com/nats-io/nats.go"
)

var storeService *service.StoreService

func init() {
	storeService = NewStoreService()
}

func NewStoreService() *service.StoreService {
	return &service.StoreService{}
}

func updateStoreLogo(message *nats.Msg) {
	var storeInfo request.FileRequest
	if err := json.Unmarshal(message.Data, &storeInfo); err != nil {
		replyData := fmt.Sprintf("error occured while trying to get message data %v", err)
		message.Respond([]byte(replyData))
		return
	}
	err := storeService.ChangeStoreLogo(storeInfo)
	if err != nil {
		replyData := fmt.Sprintf("error occured while trying to update products store %v", err)
		message.Respond([]byte(replyData))
		return
	}
	message.Respond([]byte(nil))
}

func getStoreById(message *nats.Msg) {
	var storeInfo request.GetStoreRequest
	if err := json.Unmarshal(message.Data, &storeInfo); err != nil {
		replyData := fmt.Sprintf("error occured while trying to get message data %v", err)
		message.Respond([]byte(replyData))
		return
	}
	store, err := storeService.GetStoreByID(storeInfo)
	if err != nil {
        // Convert store data to JSON if it's a valid type
        storeData, ok := store.Data.(string)
        if !ok {
            message.Respond([]byte("failed to convert store data"))
            return
        }

        message.Respond([]byte(storeData))
        return
	}
	message.Respond([]byte(nil))
}

func changeStoreCover(message *nats.Msg) {
	var storeInfo request.CoverRequest
	if err := json.Unmarshal(message.Data, &storeInfo); err != nil {
		replyData := fmt.Sprintf("error occured while trying to get message data %v", err)
		message.Respond([]byte(replyData))
		return
	}
	err := storeService.ChangeStoreCover(storeInfo)
	if err != nil {
		replyData := fmt.Sprintf("error occured while trying to update products store %v", err)
		message.Respond([]byte(replyData))
		return
	}
	message.Respond([]byte(nil))
}
