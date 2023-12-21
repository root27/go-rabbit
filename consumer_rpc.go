package amqplib

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Receive_Rpc(conn *amqp.Connection, queue string, callback func([]byte) []byte, content_type string) {

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

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
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

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		for d := range msgs {

			log.Printf("Received a message: %s", d.Body)

			response := callback(d.Body)

			err = ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   content_type,
					Body:          []byte(response),
					CorrelationId: d.CorrelationId,
				})
			if err != nil {
				log.Fatal(err)
				return
			}
			d.Ack(false)
		}

	}()

	<-forever

}
