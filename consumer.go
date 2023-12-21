package amqplib

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Receive(conn *amqp.Connection, queue string) (<-chan any, error) {

	ch, err := conn.Channel()
	if err != nil {

		log.Fatal(err)
		return nil, err
	}

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
		return nil, err
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
		return nil, err
	}

	resChan := make(chan any)

	go func() {

		for d := range msgs {

			resChan <- d.Body
		}

		conn.Close()
		ch.Close()
		close(resChan)
	}()

	return resChan, nil

}
