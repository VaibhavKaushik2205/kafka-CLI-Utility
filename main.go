package main

import (
	"fmt"

	"kafka-golang/ccloud"
	"kafka-golang/consumer"
	"kafka-golang/overview"
	"kafka-golang/producer"
	"kafka-golang/topic"
)

func main() {

	// Initialization
	configFile := ccloud.ParseArgs()
	conf := ccloud.ReadCCloudConfig(*configFile)

	// Select a choice untill you want to close the application
select_choice:
	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Printf("1. Create New Topic\n2. Create New Producer\n3. Create New Consumer\n4. Delete Topic\n5. Consumer Groups Overview\n6. Topics overview\n7. Close\n")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			topic.CreateTopic(conf)

		case 2:
			producer.CreateProducer(conf)

		case 3:
			consumer.Consume(conf)

		case 4:
			topic.DeleteTopic(conf)

		case 5:
			overview.ConsumerGroupsOverview(conf)

		case 6:
			overview.ListTopics(conf)

		case 7:
			break select_choice

		default:
			fmt.Println("Enter a valid choice...")
		}

	}
}
