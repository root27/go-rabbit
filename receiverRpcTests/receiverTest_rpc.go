package main

import (
	"log"
	"strconv"

	amqplib "github.com/root27/go-rabbit"
)

func main() {

	conn, err := amqplib.Connect("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	log.Printf("Connected to RabbitMQ")

	log.Printf("Waiting for RPC requests...")

	amqplib.Receive_Rpc(conn, "test_rpc", testFunction, "text/plain")

}

func testFunction(data []byte) []byte {

	integer, err := strconv.Atoi(string(data))

	if err != nil {
		panic(err)
	}

	return []byte(strconv.Itoa(integer * 2))

}
