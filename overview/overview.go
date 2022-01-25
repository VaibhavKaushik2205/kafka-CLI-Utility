package overview

import (
	"fmt"
	"kafka-golang/connect"
	"os/exec"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func ConsumerGroupsOverview(conf map[string]string) {

	// Get kafka path
	path := conf["kafka.dir"]

	// Describe kafka server consumers
	fmt.Printf("Servers: %v\n", conf["bootstrap.servers"])
	fmt.Printf("Path: %v\n", path)

	// Exec command to list consumer groups overview
	cmd, err := exec.Command(path, "--bootstrap-server", "localhost:9092", "--all-groups", "count_errors", "--describe").Output()

	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	output := string(cmd)
	fmt.Println((output))

}

func ListTopics(conf map[string]string) {

	// Create ConfigMap
	configMap := kafka.ConfigMap{
		"bootstrap.servers": conf["bootstrap.servers"]}

	// Create new AdminClient
	a := connect.Connect(configMap)

	metadata, err := a.GetMetadata(nil, true, 5000)

	if err != nil {
		fmt.Printf("Error get Topics MetaData: %s", err)
		return
	}

	for _, topic := range metadata.Topics {
		if topic.Topic == "__consumer_offsets" {
			continue
		}
		fmt.Printf("Topic: %v\n", topic)
	}
}
