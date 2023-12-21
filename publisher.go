package amqplib

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Send(conn *amqp.Connection, queue string, body []byte, content_type string) (string, error) {
	ch, err := conn.Channel()
	if err != nil {
		return "", err
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
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: content_type,
			Body:        []byte(body),
		})
	if err != nil {
		return "", err
	}

	return "Message sent", nil
}
