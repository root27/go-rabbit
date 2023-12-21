package amqplib

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func CloseConnection(conn *amqp.Connection) error {
	return conn.Close()
}
