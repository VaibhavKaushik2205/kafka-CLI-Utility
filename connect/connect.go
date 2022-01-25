package connect

import (
	"fmt"
	"os"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func Connect(config kafka.ConfigMap) *kafka.AdminClient {

	// Connct to brokers
	a, err := kafka.NewAdminClient(&config)
	if err != nil {
		fmt.Printf("Failed to create new admin client from producer: %s", err)
		os.Exit(1)
	}
	return a
}
