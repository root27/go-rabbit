package amqplib

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect(uri string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
