package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	nats_pkg "store-service/pkg/transport/nats"

	"github.com/nats-io/nats.go"
)

type Subscription struct {
	Subject    string `json:"subject"`
	QueueGroup string `json:"queueGroup"`
	Handler    string `json:"handler"`
}

type Subscriptions struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

var handlers = map[string]func(*nats.Msg){
	"updateStoreLogo": updateStoreLogo,
	"getStoreById": getStoreById,
	"changeStoreCover": changeStoreCover,
}

func LoadSubscriptions() {
	nc, err := nats_pkg.GetNatsInstance()
	if err != nil {
		log.Fatalf("can't connect to NATS : %v", err)
	}
	defer nc.Close()

	subscriptionsData, err := os.ReadFile("config/subscriptions.json")
	if err != nil {
		log.Fatal("Error reading configuration file:", err)
	}

	var subscriptions Subscriptions
	if err := json.Unmarshal(subscriptionsData, &subscriptions.Subscriptions); err != nil {
		log.Fatal("Error parsing configuration:", err)
	}
	for _, sub := range subscriptions.Subscriptions {
		handlerFunc, ok := handlers[sub.Handler]
		if !ok {
			log.Printf("Handler function %s not found", sub.Handler)
			continue
		}

		_, err := nc.QueueSubscribe(sub.Subject, sub.QueueGroup, handlerFunc)
		if err != nil {
			log.Printf("Error subscribing to subject %s: %v", sub.Subject, err)
		}
	}
	fmt.Printf("\n nats listen for %v subscriptions \n", len(subscriptions.Subscriptions))
	select {}
}
