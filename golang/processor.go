package main 

import (
	"fmt"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	consumer, err := getConsumer()
	if err != nil {
		fmt.Println(err)
	}
	for {
		msg, err := consumer.ReadMessage(time.Second)
		 if err == nil {
			fmt.Println(string(msg.Value))
		 } else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		 }
	}
}