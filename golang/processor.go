package main 

import (
	"syscall"
	"os"
	"fmt"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"os/signal"
)

func main() {
	//Since the topic is not created when the containers are up (using docker-composer up -d), 
	//getting a consumser (getConsumer()) returns without proper metadata.
	//Therefore, a delay is added during which we can create necessary topic
	//TODO: How do we do this production?
	time.Sleep(60 * time.Second)
	consumer, err := getConsumer()
	if err != nil {
		fmt.Printf("Failed to create consumer. %s", err)
		os.Exit(1)
	}
	
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	run := true
	for run {
		select {
		default:
			ev, err := consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				if err.(kafka.Error).IsTimeout() {
					continue
				}
				fmt.Printf("Consumer Error: %v (%v)\n", err, ev)
			}
			fmt.Println(string(ev.Key), string(ev.Value))
		case sig := <-sigchan:
			fmt.Printf("Caught Signal %v. Terminating.", sig)
			run = false
		}
	}
	consumer.Close()
}