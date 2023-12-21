package main

import (
	"encoding/json"
	"log"

	amqplib "github.com/root27/go-rabbit"
)

func main() {

	// Connect to RabbitMQ
	conn, err := amqplib.Connect("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	defer amqplib.CloseConnection(conn)

	log.Printf("connected to RabbitMQ")

	// Receive a message

	amqplib.Receive(conn, "test2", ConverttoJson)

}

var test map[string]interface{}

func ConverttoJson(body []byte) {
	err := json.Unmarshal(body, &test)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("message converted: %s", test)

}
