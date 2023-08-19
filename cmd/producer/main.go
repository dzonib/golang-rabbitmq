package main

import (
	"log"
	"time"

	"github.com/dzonib/golang-rabbitmq/internal"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("king", "kong", "localhost:5672", "customers")
	if err != nil {
		panic(err)
	}

	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	//make new rabbit client with the connection that we have ( we are reusing connection )
	client, err := internal.NewRabbitMQClient(conn)

	if err != nil {
		panic(err)
	}

	defer func(client internal.RabbitClient) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	time.Sleep(20 * time.Second)

	log.Println(client)
}
