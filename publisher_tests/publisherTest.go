package main

import (
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
	msg := "Hello World"
	_, err = amqplib.Send(conn, "test2", []byte(msg), "text/plain")
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("message sent: %s", msg)

}
