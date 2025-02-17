package nats

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

func GetNatsInstance() (*nats.Conn, error) {
	uri := os.Getenv("NATS_URI")

	if uri == "" {
		uri = nats.DefaultURL
	}
	fmt.Print(uri)
	nc, err := nats.Connect(uri)

	if err != nil {
		log.Fatalf("can't connect to NATS : %v", err)
		return nil, err
	}
	return nc, nil

}
