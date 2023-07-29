package main

import (
	"fmt"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

	con, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "broker_one:9092",
		"group.id": "gokafka",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Println(err)
	}
	
	con.SubscribeTopics([]string{"http_logs"}, nil)

	run := true
	for run {
		 msg, err := con.ReadMessage(time.Second)
		 if err == nil {
			fmt.Println(string(msg.Value))
		 } else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		 }
	}
	con.Close()
}