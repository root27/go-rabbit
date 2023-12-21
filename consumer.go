package amqplib

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Receive(conn *amqp.Connection, queue string, callback func([]byte)) {

	ch, err := conn.Channel()
	if err != nil {

		log.Fatal(err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {

		log.Fatal(err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {

		log.Fatal(err)
		return
	}

	var forever chan struct{}

	go func() {

		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			callback(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever

}
