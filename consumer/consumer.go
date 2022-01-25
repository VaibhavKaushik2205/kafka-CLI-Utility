package consumer

import (
	// "encoding/json"
	"fmt"
	"kafka-golang/ccloud"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// RecordValue represents the struct of the value in a Kafka message
type RecordValue ccloud.RecordValue

func Consume(conf map[string]string) {

	// conf := ccloud.ReadCCloudConfig(*configFile)

	fmt.Println("Enter group Id")
	var groupId string
	fmt.Scan(&groupId)

	//Topic to subscribe
	fmt.Println("Enter topic to subscribe")
	var topic string
	fmt.Scan(&topic)

	// Create Consumer instance
	totalMessagesCount := 0

	createConsumer(conf, groupId, topic, &totalMessagesCount)

}

func createConsumer(conf map[string]string, groupId string, topic string, totalMessageCount *int) *kafka.Consumer {

	configMap := kafka.ConfigMap{
		"bootstrap.servers": conf["bootstrap.servers"],
		"group.id":          groupId,
		"auto.offset.reset": "earliest"}

	// Create new Consumer
	c, err := kafka.NewConsumer(&configMap)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	// Subscribe to topic
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Unable to Subscribe consumer: %s", err)
		os.Exit(1)
	}

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	fmt.Println("Starting Consumer...")
	counsumerMessageCount := 0
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			msg, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			recordKey := string(msg.Key)
			recordValue := msg.Value
			recordValueStr := string(msg.Value)

			fmt.Printf("Key Value: %v\n", recordKey)
			fmt.Printf("Message Value (UTF-8): %v\n", recordValue)
			fmt.Printf("Message Value (String): %v\n", recordValueStr)
			counsumerMessageCount++

			// Unmarshal data if needed
			// here

			fmt.Printf("Consumed record with key %s and value %s, and updated total count to %d\n", recordKey, recordValue, counsumerMessageCount)
		}
	}

	fmt.Printf("Consumer read %v messages\n", counsumerMessageCount)
	*totalMessageCount += counsumerMessageCount
	fmt.Printf("Closing consumer\n")
	c.Close()

	return c
}
