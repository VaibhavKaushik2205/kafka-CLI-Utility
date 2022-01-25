package topic

import (
	"context"
	"fmt"
	"kafka-golang/connect"
	"os"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func CreateTopic(conf map[string]string) {

	// Create ConfigMap
	configMap := kafka.ConfigMap{
		"bootstrap.servers": conf["bootstrap.servers"]}
	// Create new AdminClient
	a := connect.Connect(configMap)

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Create topics on cluster.
	// Set Admin options to wait up to 60s for the operation to finish on the remote cluster
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Printf("ParseDuration(60s): %s", err)
		os.Exit(1)
	}

	var topicName string
	fmt.Println("New Topic Name")
	fmt.Scan(&topicName)

	var partitions int
	fmt.Println("No of partitions")
	fmt.Scan(&partitions)

	var rf int
	fmt.Println("Replication Factor")
	fmt.Scan(&rf)

	results, err := a.CreateTopics(
		ctx,
		// Multiple topics can be created simultaneously
		// by providing more TopicSpecification structs here.
		[]kafka.TopicSpecification{{
			Topic:             topicName,
			NumPartitions:     partitions,
			ReplicationFactor: rf}},
		// Admin options
		kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Admin Client request error: %v\n", err)
		os.Exit(1)
	}
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError && result.Error.Code() != kafka.ErrTopicAlreadyExists {
			fmt.Printf("Failed to create topic: %v\n", result.Error)
			os.Exit(1)
		}
		fmt.Printf("%v\n", result)
	}
	a.Close()

}

func DeleteTopic(conf map[string]string) {

	// Create ConfigMap
	configMap := kafka.ConfigMap{
		"bootstrap.servers": conf["bootstrap.servers"]}
	// Create new AdminClient
	a := connect.Connect(configMap)

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var topic string
	fmt.Println("Select topic to delete")
	fmt.Scan(&topic)
	_, err := a.DeleteTopics(ctx, []string{topic})

	if err != nil {
		fmt.Printf("Failed to Delete topic: %s", err)
		os.Exit(1)
	}

	a.Close()

}
