package ccloud

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// RecordValue represents the struct of the value in a Kafka message
type RecordValue struct {
	Count int
}

var configPath string = "confluent/librdkafka.config"

// ParseArgs parses the command line arguments and
// returns the config file and topic on success, or exits on error
func ParseArgs() *string {

	configFile := flag.String("f", configPath, "Path to Confluent Cloud configuration file")
	flag.Parse()
	if *configFile == "" {
		flag.Usage()
		os.Exit(2) // the same exit code flag.Parse uses
	}

	return configFile

}

// ReadCCloudConfig reads the file specified by configFile and
// creates a map of key-value pairs that correspond to each
// line of the file. ReadCCloudConfig returns the map on success,
// or exits on error
func ReadCCloudConfig(configFile string) map[string]string {

	m := make(map[string]string)

	file, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	return m

}
