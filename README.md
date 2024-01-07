# Rabbitmq Package for Golang

It includes basic operations for Rabbitmq. These operations are:

- Basic Publish/Consumer
- Publish/Consumer RPC


## Installation

```bash

go get github.com/root27/go-rabbit

```

## Usage

Some usage examples are given below.

### Basic Publish/Consumer

- Basic Publish


```go

package main

import (
	"encoding/json"
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
	msg := SendMessage{"message": "Hello World!"}

	byteData, err := json.Marshal(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = amqplib.Send(conn, "test2", []byte(byteData), "application/json")
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("message sent: %s", msg)

}

type SendMessage map[string]interface{}

```

- Basic Consumer


```go

package main

import (
	"log"

	amqplib "github.com/root27/go-rabbit"
)

func main() {

	// Connect to RabbitMQ
	conn, err := amqplib.Connect("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("connected to RabbitMQ")

	log.Printf("Messages waiting...")

	// Receive a message

	msgs, err := amqplib.Receive(conn, "test2")  // return a channel

	if err != nil {
		log.Fatal(err)
	}


    // msgs is a channel, so we can use it in a for loop


	for d := range msgs {

		log.Printf("received a message: %s", d)

		log.Printf("Type of data received: %T", d)
	}

}

```

### Publish/Consumer RPC

- Publish RPC

```go

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


    // Send RPC accepts 4 parameters: connection, queue name, data, data type

    // Returns consume RPC response

	response, err := amqplib.Send_Rpc(conn, "test_rpc", []byte(testData), "text/plain")

	if err != nil {
		panic(err)
	}

	log.Printf("Response: %s", response)

}

```

- Consumer RPC

```go

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

    // Receive RPC accepts 4 parameters: connection, queue name, function, data type


    // Function must be a function that accepts a byte array and returns a byte array


	amqplib.Receive_Rpc(conn, "test_rpc", testFunction, "text/plain")

}

func testFunction(data []byte) []byte {

	integer, err := strconv.Atoi(string(data))

	if err != nil {
		panic(err)
	}

	return []byte(strconv.Itoa(integer * 2))

}

``````









