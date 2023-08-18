package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func getConsumer() (*kafka.Consumer, error) {
	con, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "broker_one:9092",
		"group.id": "gokafka",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	con.SubscribeTopics([]string{"http_logs"}, nil)
	return con, nil
}