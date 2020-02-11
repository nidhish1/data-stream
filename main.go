package main
import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	//"sync"
	//https://gist.github.com/savaki/a19dcc1e72cb5d621118fbee1db4e61f
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka",
		//key.deserializer
		//value.deserializer
		"group.id": "Consumer1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	//var wg sync.WaitGroup
	c.SubscribeTopics([]string{"server1.dbo.Lynk_DriverAvailable_AfterOffline"}, nil)
	//c.SubscribeTopics([]string{"server1.dbo.* "}, nil)   ----- to subscribe to all the topics

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			
		} else {
			
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	 // wg.wait()
	c.Close()
}

// func Lynk_DriverAvailable_AfterOffline (c *Message) {
// 	if (msg.TopicPar {

// 	}
//  }