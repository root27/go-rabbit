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

	testData := strconv.Itoa(5)

	response, err := amqplib.Send_Rpc(conn, "test_rpc", []byte(testData), "text/plain")

	if err != nil {
		panic(err)
	}

	log.Printf("Response: %s", response)

}
