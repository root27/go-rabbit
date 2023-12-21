package main

import (
	"log"

	amqplib "github.com/root27/go-rabbit"
)

func main() {

	// Connect to RabbitMQ
	conn, err := amqplib.Connect("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("connected to RabbitMQ")

	log.Printf("Messages waiting...")

	// Receive a message

	msgs, err := amqplib.Receive(conn, "test2")

	if err != nil {
		log.Fatal(err)
	}

	for d := range msgs {

		log.Printf("received a message: %s", d)

		log.Printf("Type of data received: %T", d)
	}

}
