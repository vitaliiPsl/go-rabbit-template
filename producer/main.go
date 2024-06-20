package main

import (
	"encoding/json"
	"log"
	"producer/config"
	"producer/messaging"
)

func main() {
	config := config.ReadConfig()

	connectionManager := messaging.NewRabbitConnectionManager(config)
	rabbitClient := messaging.NewRabbitClient(connectionManager)

	payload := map[string]string{"id": "je1pot31", "email": "j.doe@mail.com", "message": "Ce type là-bas n'est pas réel"}
	payloadJson, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Fatalf("Error while marshaling json: err=[%v]\n", err.Error())
	}

	log.Printf("Payload=[%v]\n", string(payloadJson))
	rabbitClient.Publish("message.rk", "message.exchange", string(payloadJson))
}
