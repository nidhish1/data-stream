package repository

import (
	"fmt"
	//"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	//"github.com/spf13/viper"
)

func ReceiveFromKafka() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka",
		"group.id": "consumer1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	c.SubscribeTopics([]string{"server1.dbo.Lynk_DriverAvailable_AfterOffline"}, nil)

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Received from Kafka %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()

}
