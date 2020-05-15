package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaServer, kafkaTopic string

// const groupID = "confluent_cli_consumer_21283b45-cd87-443c-9d85-99618c0bd402"

func init() {
	kafkaServer = readFromENV("KAFKA_BROKER", "pkc-ldvr1.asia-southeast1.gcp.confluent.cloud:9092")
	kafkaTopic = readFromENV("KAFKA_TOPIC", "test")

	fmt.Println("Kafka Broker - ", kafkaServer)
	fmt.Println("Kafka topic - ", kafkaTopic)
}

func main() {

	c1, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "pkc-ldvr1.asia-southeast1.gcp.confluent.cloud:9092",
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     "J455WLJIMGCBRZEO",
		"sasl.password":     "sH8SgRiEy3ManX9R0HT7CTDWtcB0m1s3oZJ6lnlSMWUBydjHEpx+d+r45qL/l2CT",
		// "group.id":          "confluent_cli_consumer_21283b45-cd87-443c-9d85-99618c0bd402",
		"group.id":          "test-consumer-group",
		"auto.offset.reset": "earliest"})

	if err != nil {
		panic(err)
	}

	c1.Subscribe(kafkaTopic, nil)
	for {
		fmt.Println("c1 : ")
		msg, err := c1.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
		fmt.Println("End >>>>> c1")
	}

	c1.Close()

}

func readFromENV(key, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}
