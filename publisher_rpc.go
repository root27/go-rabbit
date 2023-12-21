package amqplib

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Send_Rpc(conn *amqp.Connection, queue string, data []byte, content_type string) ([]byte, error) {

	ch, err := conn.Channel()
	if err != nil {
		return []byte(""), err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return []byte(""), err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return []byte(""), err
	}

	corrId := randomString(32)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		queue,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:   content_type,
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(data),
		})

	if err != nil {
		return []byte(""), err
	}

	for d := range msgs {
		if corrId == d.CorrelationId {

			return d.Body, nil

		}
	}

	return []byte(""), nil

}
