package messaging

import (
	"consumer/config"
	"github.com/streadway/amqp"
	"log"
)

type RabbitConnectionManager struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitConnectionManager(cfg *config.Config) *RabbitConnectionManager {
	conn, err := amqp.Dial(cfg.RABBIT_URL)
	if err != nil {
		log.Fatalf("Failed to open connection. Err=[%v]", err.Error())
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel. Err=[%v]", err.Error())
	}

	err = channel.ExchangeDeclare(
		cfg.RABBIT_MAIN_EXCHANGE_NAME,
		cfg.RABBIT_MAIN_EXCHANGE_TYPE,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to create exchange. Err=[%v]", err.Error())
	}

	_, err = channel.QueueDeclare(
		cfg.RABBIT_MAIN_QUEUE_NAME,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to create queue. Err=[%v]", err.Error())
	}

	err = channel.QueueBind(
		cfg.RABBIT_MAIN_QUEUE_NAME,
		cfg.RABBIT_MAIN_ROUTING_KEY,
		cfg.RABBIT_MAIN_EXCHANGE_NAME,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue to exchange. Err=[%v]", err.Error())
	}

	return &RabbitConnectionManager{
		Connection: conn,
		Channel:    channel,
	}
}

func (cm *RabbitConnectionManager) Close() {
	err := cm.Channel.Close()
	if err != nil {
		log.Printf("Failed to close RabbitMQ channel. Err=[%v]", err.Error())
	}
	err = cm.Connection.Close()
	if err != nil {
		log.Printf("Failed to close RabbitMQ connection. Err=[%v]", err.Error())
	}
	log.Println("RabbitMQ connection and channel closed.")
}
