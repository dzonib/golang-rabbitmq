package internal

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitClient is wrapper around official amqp client in which we will add additional functionality to
type RabbitClient struct {
	// the connection used by the client
	// connection in rabbitmq is tcp connection
	// good rule is to reuse the connection in whole app and span channels on concurrent tasks that are running
	conn *amqp.Connection

	// channel is multiplex sub connection
	// used to process and send messages
	ch *amqp.Channel
}

func ConnectRabbitMQ(username, password, host, vhost string) (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, host, vhost))
}

func NewRabbitMQClient(conn *amqp.Connection) (RabbitClient, error) {
	ch, err := conn.Channel()

	if err != nil {
		return RabbitClient{}, err
	}

	return RabbitClient{conn: conn, ch: ch}, nil
}

func (rn RabbitClient) Close() error {
	return rn.ch.Close()
}
