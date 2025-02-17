package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

func GetNatsInstance() (*nats.Conn, error) {
	// NATS server initialization
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Error connecting to NATS server: %v", err)
		return nil, err
	}
	return nc, nil
}