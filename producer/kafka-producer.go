package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaServer, kafkaTopic string

func init() {
	kafkaServer = readFromENV("KAFKA_BROKER", "pkc-ldvr1.asia-southeast1.gcp.confluent.cloud:9092")
	kafkaTopic = readFromENV("KAFKA_TOPIC", "test")

	fmt.Println("Kafka Broker - ", kafkaServer)
	fmt.Println("Kafka topic - ", kafkaTopic)
}
func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     "J455WLJIMGCBRZEO",
		"sasl.password":     "sH8SgRiEy3ManX9R0HT7CTDWtcB0m1s3oZJ6lnlSMWUBydjHEpx+d+r45qL/l2CT",
		"acks":              "all"})
	if err != nil {
		panic(err)
	}
	defer p.Close()
	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	// Produce messages to topic (asynchronously)
	topic := "test"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)

	// if err != nil {
	// 	fmt.Printf("Failed to create producer: %s\n", err)
	// 	os.Exit(1)
	// }

	// topic := "test"

	// deliveryChan := make(chan kafka.Event, 10000)
	// err = p.Produce(&kafka.Message{
	// 	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 	Value:          []byte("value eiei")},
	// 	deliveryChan,
	// )

	// e := <-deliveryChan
	// m := e.(*kafka.Message)

	// if m.TopicPartition.Error != nil {
	// 	fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	// } else {
	// 	fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
	// 		*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	// }

	// close(deliveryChan)
	// if err != nil {
	// 	panic(err)
	// }

	// defer p.Close()

	// // Delivery report handler for produced messages
	// go func() {
	// 	for e := range p.Events() {
	// 		switch ev := e.(type) {
	// 		case *kafka.Message:
	// 			if ev.TopicPartition.Error != nil {
	// 				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	// 			} else {
	// 				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
	// 			}
	// 		}
	// 	}
	// }()

	// // Produce messages to topic (asynchronously)
	// topic := kafkaTopic
	// for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
	// 	p.Produce(&kafka.Message{
	// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 		Value:          []byte(word),
	// 	}, nil)
	// }

	// // Wait for message deliveries before shutting down
	// p.Flush(15 * 1000)
	// }

	// producer, producerCreateErr := kafka.NewProducer(&kafka.ConfigMap{
	// 	"bootstrap.servers": kafkaServer,
	// 	"security.protocol": "SASL_SSL",
	// 	"sasl.mechanisms":   "PLAIN",
	// 	"sasl.username":     "J455WLJIMGCBRZEO",
	// 	"sasl.password":     "sH8SgRiEy3ManX9R0HT7CTDWtcB0m1s3oZJ6lnlSMWUBydjHEpx+d+r45qL/l2CT",
	// 	"acks":              "all"})

	// if producerCreateErr != nil {
	// 	fmt.Println("Failed to create producer due to ", producerCreateErr)
	// 	os.Exit(1)
	// }
	// for {
	// 	for i := 0; i < 5; i++ {
	// 		value := time.Now().String()
	// 		producerErr := producer.Produce(&kafka.Message{
	// 			TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
	// 			Value:          []byte(value),
	// 		}, nil)

	// 		if producerErr != nil {
	// 			fmt.Println("unable to enqueue message ", value)
	// 		}
	// 		event := <-producer.Events()

	// 		fmt.Println(event)

	// 		message := event.(*kafka.Message)

	// 		if message.TopicPartition.Error != nil {
	// 			fmt.Println("Delivery failed due to error ", message.TopicPartition.Error)
	// 		} else {
	// 			fmt.Println("Delivered message to offset " + message.TopicPartition.Offset.String() + " in partition " + message.TopicPartition.String())
	// 		}
	// 	}

	// 	time.Sleep(time.Second * 5) //wait for 5 seconds before sending another batch of messages
	// }

}
func readFromENV(key, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}
