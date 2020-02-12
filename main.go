package main
import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"github.com/nidhish1/data-stream/repository"
	//"sync"
	//https://gist.github.com/savaki/a19dcc1e72cb5d621118fbee1db4e61f
)

func main() {
	repository.ReceiveFromKafka()
}

// func Lynk_DriverAvailable_AfterOffline (c *Message) {
// 	if (msg.TopicPar {

// 	}
//  }