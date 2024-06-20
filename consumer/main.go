package main

import (
	"consumer/config"
	"consumer/messaging"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	config := config.ReadConfig()

	connectionManager := messaging.NewRabbitConnectionManager(config)
	rabbitClient := messaging.NewRabbitClient(connectionManager)

	err := rabbitClient.Consume("message.queue", func(d amqp.Delivery) {
		log.Printf("Received a message: %s", d.Body)
	})
	if err != nil {
		log.Fatalf("Failed to start consuming messages: %s", err)
	}

	forever := make(chan int, 1)
	<- forever
}
