package messaging

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitClient struct {
	Manager *RabbitConnectionManager
}

func NewRabbitClient(connectionManager *RabbitConnectionManager) *RabbitClient {
	return &RabbitClient{Manager: connectionManager}
}

func (client *RabbitClient) Publish(routingKey, exchange, message string) error {
	payload := amqp.Publishing{
		ContentType: "text/json",
		Body:        []byte(message),
	}

	return client.Manager.Channel.Publish(exchange, routingKey, false, false, payload)
}

func (client *RabbitClient) Consume(queue string, handler func(amqp.Delivery)) error {
	msgs, err := client.Manager.Channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("Failed to start consuming messages from queue. Queue=[%v], Err=[%v]", queue, err.Error())
		return err
	}

	go func() {
		for d := range msgs {
			handler(d)
		}
	}()
	return nil
}

func (client *RabbitClient) Close() {
	client.Manager.Close()
}
