package main

import (
	"encoding/json"
	"fmt"
	"log"

	amqplib "github.com/root27/go-rabbit"
)

func main() {

	// Connect to RabbitMQ
	conn, err := amqplib.Connect("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	log.Printf("connected to RabbitMQ")

	// Send a message
	msg := SendMessage{"message": "Hello World!"}

	byteData, err := json.Marshal(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = amqplib.Send(conn, "test2", []byte(byteData), "application/json")
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("message sent: %s", msg)

}

type SendMessage map[string]interface{}
