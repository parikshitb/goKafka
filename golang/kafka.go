package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func GetConsumer() (*kafka.Consumer, error) {
	configurations := getConsumerConfig()
	con, err := kafka.NewConsumer(&configurations)
	if err != nil {
		return nil, err
	}

	con.SubscribeTopics([]string{"http_logs"}, nil)
	return con, nil
}

func getConsumerConfig() kafka.ConfigMap {
	configurations := make(map[string]kafka.ConfigValue)
	configurations["bootstrap.servers"] = "broker_one:9092"
	configurations["group.id"] = "gokafka"
	configurations["auto.offset.reset"] = "earliest"
	return configurations
}
