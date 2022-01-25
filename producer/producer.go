package producer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"kafka-golang/ccloud"
	"os"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// RecordValue represents the struct of the value in a Kafka message
type RecordValue ccloud.RecordValue

func CreateProducer(conf map[string]string) {

	// Create ConfigMap
	configMap := kafka.ConfigMap{
		"bootstrap.servers": conf["bootstrap.servers"]}

	// Select topic
	var topicName string
	fmt.Println("Enter Topic to produce messages to")
	fmt.Scan(&topicName)

	// Create Producer instance
	p, err := kafka.NewProducer(&configMap)
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	messageCount := 0
	for {
		var sendMessage string
		fmt.Println("Do you want to send a new message? (yes/no)")
		fmt.Scan(&sendMessage)
		if sendMessage == "no" {
			break
		} else if sendMessage == "yes" {
			sendNewMessage(p, topicName)
			messageCount++
		} else {
			fmt.Println("Not a valid choice. Choose again")
		}
	}

	// Wait for all messages to be delivered
	p.Flush(messageCount * 1000)

	fmt.Printf("%v messages produced to topic %s!", messageCount, topicName)
	p.Close()
}

func sendNewMessage(p *kafka.Producer, topicName string) {

	//Input Stream
	in := bufio.NewReader(os.Stdin)

	var recordKey string = ""
	fmt.Println("Do you want to provide a key? (yes/no)")
	var key string
	fmt.Scan(&key)
	if key == "yes" {
		fmt.Println("Enter Key")
		fmt.Scan(&recordKey)
	}

	fmt.Println("Enter message")
	message, _ := in.ReadString('\n')

	recordValue, _ := json.Marshal(&message)

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny},
		Key:            []byte(recordKey),
		Value:          []byte(recordValue),
	}, nil)

	fmt.Printf("Message produced to topic %s!\n", topicName)
}
